#!/bin/bash
# test Issue #33
readonly DIR=$(dirname $0)
source $DIR/env.txt

${MAILSEND} \
        -debug \
        -subject "this is a test" \
        -from "noreply@example.com" \
        -to "${TO}" \
        -smtp "$SMTP_SERVER" \
        -cs "ISO-8859-1" \
        auth \
        -user "$FROM" \
        -pass "$PASS" \
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
