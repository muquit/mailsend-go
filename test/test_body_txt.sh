#!/bin/bash

########################################################################
# Test STMP info
# Part of mailsend-go muquit@muquit.com 
########################################################################
#
readonly DIR=$(dirname $0)
source ${DIR}/env.txt

$MAILSEND -smtp $SMTP_SERVER -port $TLS_PORT \
    -debug \
    -sub "Mail with a Text body" \
    -t "${TO}" \
    -f "${FROM}" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -msg 'This is a text message in mail body'
