#!/bin/bash

readonly DIR=$(dirname $0)
source ${DIR}/env.txt

${MAILSEND}  \
        -debug \
        -subject "¡Hola, señor!" \
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
            -file "$DIR/cat.jpg" \
            -inline
