#!/bin/bash

readonly DIR=$(dirname $0)
source $DIR/env.txt

${MAILSEND} \
        -debug \
        -subject "this is a test using SSL" \
        -from "${FROM}" \
        -to "${TO}" \
        -smtp "$SMTP_SERVER" \
        -port $SSL_PORT \
        -ssl \
        auth \
        -user "$FROM" \
        -pass "$SMTP_USRE_PASS" \
        body \
        -msg "তোমাকে ভালোবাসি বোলে আজো স্বপ্ন দেখি <br>Cats attached, flower is inline" \
        -mime-type "text/html" \
        attach \
            -file "$HOME/mailsend-data/cats.jpg" \
            -mime-type "image/jpeg"  \
         attach \
            -file "$HOME/mailsend-data/flower.jpg" \
            -mime-type "image/gif"  \
            -inline
