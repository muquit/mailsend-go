#!/bin/bash

readonly DIR=$(dirname $0)
source $DIR/env.txt

${MAILSEND} \
        -debug \
        -subject "this is a test" \
        -fname "Mailsend Real" \
        -from "${FROM}" \
        -to "${TO}" \
        -smtp "$SMTP_SERVER" \
        -port $TLS_PORT \
        auth \
        -user "$FROM" \
        -pass "$SMTP_USER_PASS" \
        body \
        -msg "This is a message" \
        attach \
            -file "$HOME/Downloads/cat.jpg" \
            -mime-type "image/jpeg"  \
            -inline
