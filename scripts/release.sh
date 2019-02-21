#!/bin/bash

# Wrapper for golrealse
# Usage: release.sh   - dryrun
#        rlease.sh ok - really rlease
# muquit@muquit.com Jan-20-2019 
ARGC=$#

TF=$(mktemp)
create_chl() {
    v=$(git describe --abbrev=0 --tags)
    tag=$(git describe --abbrev=0 --tags|sed -e 's/\.//g')
    echo "Plese look at [ChangeLog](ChangeLog.md#$tag) for changes in ${v}" > $TF
}
create_chl

/bin/rm -rf ./dist
if [[ $ARGC == 1 && $1 == "ok" ]]; then
    echo "Publish..."
    : ${GITHUB_TOKEN:?"Need to set GITHUB_TOKEN"}
    goreleaser --release-notes=${TF}
elif [[ $ARGC == 1 && $1 == "snap" ]]; then
    goreleaser --snapshot --skip-publish --rm-dist --release-notes=${TF}
else
    echo "Dryrun.."
    goreleaser release --skip-publish
fi
/bin/rm -f $TF
