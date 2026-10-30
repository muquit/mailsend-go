#!/bin/bash

########################################################################
# Show the winget status of a mailsend-go release: which versions are
# published in microsoft/winget-pkgs, the PR for the given version, its
# labels, the validation pipeline checks, and what (if anything) needs
# to be done next.
#
# Requires:
#   - gh (GitHub CLI), authenticated (gh auth login) or GH_TOKEN /
#     GITHUB_TOKEN exported. A classic PAT with public_repo scope is
#     enough (no read:org needed).
#
# Usage: winget_status.sh [-p pr_number] [version]
#        version defaults to the contents of VERSION (e.g. v1.0.14)
#        -p skips the PR search and uses the given PR number
#
# Exit codes:
#   0  merged into winget-pkgs
#   1  error (no gh, no PR found, API failure)
#   2  open, nothing to do but wait (validation running or moderator queue)
#   3  open, needs action from you (CLA, author feedback, validation error)
#   4  closed without merge
#
# By Claude Code, Oct-30-2026
########################################################################

PACKAGE_ID="muquit.mailsend-go"
RELEASE_REPO="muquit/mailsend-go"
UPSTREAM="microsoft/winget-pkgs"
MANIFEST_DIR="manifests/m/muquit/mailsend-go"
VALIDATOR_APP="wingetvalidator-prod"

if ! command -v gh >/dev/null 2>&1; then
    echo "gh (GitHub CLI) is not installed" >&2
    exit 1
fi

PR_OVERRIDE=""
if [[ "$1" == "-p" ]]; then
    PR_OVERRIDE="${2#\#}"
    shift 2
fi

VERSION="$1"
if [[ -z "$VERSION" ]]; then
    if [[ ! -f VERSION ]]; then
        echo "Usage: $0 [version]  (no VERSION file found)" >&2
        exit 1
    fi
    VERSION=$(cat VERSION | tr -d '\n')
fi
# winget versions have no leading 'v'
VERSION="${VERSION#v}"

echo "winget status for ${PACKAGE_ID} ${VERSION}"
echo "========================================================"

#-----------------------------------------------------------------------
# Versions already merged into winget-pkgs
#-----------------------------------------------------------------------
published=$(gh api "repos/${UPSTREAM}/contents/${MANIFEST_DIR}" \
    --jq '.[] | select(.type=="dir") | .name' 2>/dev/null | sort -V)
if [[ -n "$published" ]]; then
    echo "Published versions : $(echo $published)"
else
    echo "Published versions : (could not list ${MANIFEST_DIR})"
fi

#-----------------------------------------------------------------------
# Find the PR for this version (newest first)
#-----------------------------------------------------------------------
if [[ -n "$PR_OVERRIDE" ]]; then
    PR="$PR_OVERRIDE"
else
    # same search that 'gh pr list' does by hand; pick the newest PR
    # whose title ends with exactly this version
    PR=$(gh pr list -R "${UPSTREAM}" --state all \
        --search "${PACKAGE_ID} ${VERSION} in:title" \
        --json number,title,createdAt \
        --jq "[.[] | select(.title | endswith(\"version ${VERSION}\"))] | sort_by(.createdAt) | last | .number // empty")
fi

if [[ -z "$PR" ]]; then
    echo
    if grep -qx "$VERSION" <<< "$published"; then
        echo "No PR found, but ${VERSION} is published. Nothing to do."
        exit 0
    fi
    echo "No PR found for ${VERSION}."
    echo "Did the Publish to Winget action run? Check with:"
    echo "  gh run list -R ${RELEASE_REPO} --workflow winget.yml -L 3"
    echo "To submit by hand: ./komac_winget.sh v${VERSION}"
    echo "If you know the PR number: $0 -p <number> v${VERSION}"
    exit 1
fi

info=$(gh api "repos/${UPSTREAM}/pulls/${PR}" \
    --jq '"\(.state)|\(.merged_at // "")|\(.created_at)|\(.html_url)|\(.head.sha)|\([.labels[].name] | join(","))"')
if [[ -z "$info" ]]; then
    echo "Could not read PR #${PR}" >&2
    exit 1
fi
IFS='|' read -r state merged_at created url sha labels <<< "$info"

echo "PR                 : #${PR}  ${url}"
echo "Created            : ${created}"
echo "State              : ${state}"
[[ -n "$merged_at" ]] && echo "Merged at          : ${merged_at}"
echo "Labels             : ${labels:-(none)}"

#-----------------------------------------------------------------------
# Validation pipeline checks
#-----------------------------------------------------------------------
checks=$(gh api "repos/${UPSTREAM}/commits/${sha}/check-runs?per_page=100" \
    --jq ".check_runs[] | select(.app.slug==\"${VALIDATOR_APP}\") | \"\(.name)|\(.conclusion // .status)\"" \
    2>/dev/null | sort)

echo
echo "Validation checks:"
if [[ -z "$checks" ]]; then
    echo "  (none reported yet, validation is queued or running)"
else
    while IFS='|' read -r name result; do
        case "$result" in
            success) mark="ok  " ;;
            neutral) mark="WARN" ;;
            failure) mark="FAIL" ;;
            *)       mark="..  " ;;
        esac
        printf "  [%s] %s (%s)\n" "$mark" "$name" "$result"
    done <<< "$checks"
fi

validation_done=0
grep -q "Validation Completed|success" <<< "$checks" && validation_done=1
install_neutral=0
grep -q "Installation Validation|neutral" <<< "$checks" && install_neutral=1

#-----------------------------------------------------------------------
# Last comment on the PR
#-----------------------------------------------------------------------
last_comment=$(gh api "repos/${UPSTREAM}/issues/${PR}/comments?per_page=100" \
    --jq 'if length == 0 then empty else last | "\(.user.login) (\(.created_at)):\n\(.body)" end' \
    2>/dev/null | grep -v '^<!--' | head -8)
if [[ -n "$last_comment" ]]; then
    echo
    echo "Last comment:"
    sed 's/^/  /' <<< "$last_comment"
fi

#-----------------------------------------------------------------------
# What it means
#-----------------------------------------------------------------------
echo
echo "Verdict:"

if [[ -n "$merged_at" ]]; then
    echo "  Merged. The winget index usually picks it up within a few hours."
    echo "  On Windows:  winget source update; winget upgrade ${PACKAGE_ID}"
    exit 0
fi

if [[ "$state" == "closed" ]]; then
    echo "  Closed without merge. Read the comments: ${url}"
    exit 4
fi

needs_action=0
IFS=',' read -r -a label_list <<< "$labels"
for label in "${label_list[@]}"; do
    case "$label" in
        Needs-CLA)
            echo "  Needs-CLA: sign the CLA by commenting on the PR:"
            echo "    gh pr comment ${PR} -R ${UPSTREAM} --body \"@microsoft-github-policy-service agree\""
            needs_action=1
            ;;
        Needs-Author-Feedback|Needs-Attention)
            echo "  ${label}: a moderator or bot is waiting on you. Read: ${url}"
            needs_action=1
            ;;
        Validation-Executable-Error)
            if [[ $validation_done -eq 1 && $install_neutral -eq 1 ]]; then
                echo "  Validation-Executable-Error is only a warning (Installation"
                echo "  Validation = neutral, Validation Completed = success)."
                echo "  Common for portable CLI tools. A moderator has to review it."
            else
                echo "  Validation-Executable-Error: installation test failed. Read: ${url}"
                needs_action=1
            fi
            ;;
        Validation-*Error*|*-Error)
            echo "  ${label}: validation failed. Read: ${url}"
            needs_action=1
            ;;
        Moderator-Approved)
            echo "  Moderator-Approved: should be merged soon."
            ;;
    esac
done

if [[ $needs_action -eq 1 ]]; then
    exit 3
fi

if [[ $validation_done -eq 1 ]]; then
    echo "  Validation finished. Waiting in the moderator queue (often a few days)."
else
    echo "  Validation not finished yet. Check again later."
fi
exit 2
