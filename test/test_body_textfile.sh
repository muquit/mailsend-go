#!/bin/bash

########################################################################
# Test STMP info
# Part of mailsend-go muquit@muquit.com 
########################################################################

readonly DIR=$(dirname $0)
source ${DIR}/env.txt

$MAILSEND \
    -debug \
    -sub "Mail with a Text bodfy from a text file" \
    -t "${TO}" \
    -f "${FROM}" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -file "$DIR/file.txt" \
    -smtp $SMTP_SERVER -port $TLS_PORT \
