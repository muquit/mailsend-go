#!/bin/bash
# Swiched to use gomail from my forked repo.
# clean things up first
# muquit@muquit.com Feb-14-2025 
echo "> clean modcahe ..."
go clean -modcache
echo "> clean cache ..."
go clean -cache
go clean -i -r
go env GOMODCACHE

mod_dir=${GOPATH}/pkg/mod
if [[ -d ${mod_dir} ]]; then
    chmod -R 777 ${mod_dir}
    echo "> clean dir ${mod_dir} ..."
    /bin/rm -rf ${mod_dir}
fi
echo "> clean go.sum ..."
/bin/rm -f go.sum
echo "> clean go.mod ..."
/bin/rm -f go.mod
echo "> initialize go.mod .."
go mod init github.com/muquit/mailsend-go
go mod tidy
echo "> List go.mod ..."
cat go.mod
echo "> List go.sum ..."
cat go.sum
