#!/bin/bash

# for v1.0.13
# use app password
# Also test Issue #76
# Sep-24-2026 
readonly DIR=$(dirname $0)
source $DIR/env.txt
SMTP_SERVER='smtp.fastmail.com'

${MAILSEND} \
        -verbose \
        -log /tmp/mailsend-go.log \
        -subject "this is a mailsend-go test" \
        -from "${FROM}" \
        -to "${TO}" \
        -smtp "$SMTP_SERVER" \
        -port $TLS_PORT \
        auth \
        -user "$FROM" \
        -pass "$PASS" \
        -verifyCert \
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
