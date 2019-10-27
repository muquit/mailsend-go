#!/bin/bash

readonly DIR=$(dirname $0)
source $DIR/env.txt

${MAILSEND} \
        -log /tmp/mailsend-go.log \
        -subject "Testing Big5 Charset" \
        -from "noreply@example.com" \
        -to "${TO}" \
        -smtp "$SMTP_SERVER" \
        -cs "Big5" \
        -port $TLS_PORT \
        auth \
        -user "$FROM" \
        -pass "$PASS" \
        body \
        -msg "中文測試"
