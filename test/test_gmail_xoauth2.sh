#!/bin/bash

########################################################################
# Test sending mail via XOAUTH2
# Aug-24-2025 
########################################################################
readonly DIR=$(dirname $0)

MAILSEND=./mailsend-go
TO=${TO}
SMTP_SERVER="${SMTP_SERVER:-smtp.gmail.com}"
TLS_PORT="${TLS_PORT:-587}"
SSL_PORT="${SSL_PORT:-465}"
FRROM=${FROM}
SMTP_USER=$FROM
: SMTP_OAUTH_TOKEN:=${SMTP_OAUTH_TOKEN}

abortTest() {
    local -r MSG=$1
    echo "${MSG}"
    exit 1
}

[[ $TO ]] || abortTest "ERROR: Set environment variable TO"
[[ $FROM ]] || abortTest "ERRROR: Set environment variable FROM"
[[ $SMTP_OAUTH_TOKEN ]] || abortTest "ERROR: Set environment variable SMTP_OAUTH_TOKEN"

${MAILSEND} \
        -log /tmp/mailsend-go.log \
        -subject "this is a test" \
        -from "noreply@example.com" \
        -to "${TO}" \
        -smtp "$SMTP_SERVER" \
        -cs "ISO-8859-1" \
        -port $TLS_PORT \
        auth \
            -user "$FROM" \
            -oauth2 \
        body \
        -msg "Cats attached, flower is inline" \
        -mime-type "text/html" \
        attach \
            -file "$HOME/mailsend-data/cats.jpg" \
            -mime-type "image/jpeg"  \
         attach \
            -file "$HOME/mailsend-data/flower.jpg" \
            -mime-type "image/gif"  \
            -inline
