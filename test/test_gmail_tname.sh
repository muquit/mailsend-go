#!/bin/bash

readonly DIR=$(dirname $0)
source $DIR/env.txt

$MAILSEND \
        -debug \
        -subject "Testing -fname and -tname" \
        -from "${FROM}" \
        -fname "Mailsend Tester" \
        -tname "John Snow" \
        -to "$TO" \
        -smtp "$SMTP_SERVER" \
        -port $TLS_PORT \
        auth \
        -user "$FROM" \
        -pass "$SMTP_USER_PASS" \
        body \
        -msg "<b>Testing -tname and -fname</b>" \
        -mime-type "text/html" \
