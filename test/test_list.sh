#!/bin/bash

readonly DIR=$(dirname $0)
source ${DIR}/env.txt

${MAILSEND} \
        -debug \
        -subject "Test sending mail to a list of users" \
        -from "${FROM}" \
        -to "${TO}" \
        -smtp "$SMTP_SERVER" \
        -port $TLS_PORT \
        auth \
        -user "$FROM" \
        -pass "$SMTP_USER_PASS" \
        body \
        -msg "Cat attached, flower is inline" \
        -mime-type "text/html" \
        attach \
            -file "$DIR/cat.jpg" \
         attach \
            -file "$DIR/flower.jpg" \
            -inline \
         -list "$HOME/mailsend-data/list.txt"
