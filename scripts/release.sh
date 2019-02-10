#!/bin/bash

# Wrapper for golrealse
# Usage: release.sh   - dryrun
#        rlease.sh ok - really rlease
# muquit@muquit.com Jan-20-2019 
ARGC=$#

/bin/rm -rf ./dist
if [[ $ARGC == 1 && $1 == "ok" ]]; then
    echo "Publish..."
    : ${GITHUB_TOKEN:?"Need to set GITHUB_TOKEN"}
    goreleaser
elif [[ $ARGC == 1 && $1 == "snap" ]]; then
    goreleaser --snapshot --skip-publish --rm-dist
else
    echo "Dryrun.."
    goreleaser release --skip-publish
fi
