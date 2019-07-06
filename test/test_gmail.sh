#!/bin/bash

readonly DIR=$(dirname $0)
source $DIR/env.txt

${MAILSEND} \
        -log /tmp/mailsend-go.log \
        -subject "this is a test" \
        -from "${FROM}" \
        -to "${TO}" \
        -smtp "$SMTP_SERVER" \
        -port $TLS_PORT \
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
