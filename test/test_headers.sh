#!/bin/bash

readonly DIR=$(dirname $0)
source $DIR/env.txt

${MAILSEND} \
        -debug \
        -subject "this is a test" \
        -fname "mailsend-go" \
        -from "${FROM}" \
        -to "${TO}" \
        -smtp "$SMTP_SERVER" \
        -port $TLS_PORT \
        auth \
        -user "$FROM" \
        -pass "$PASS" \
        body \
        -msg "Cats image's Content-Disposition is \"inline\". There are 3 custom headers." \
        attach \
            -file "$HOME/mailsend-data/cats.jpg" \
            -mime-type "image/jpeg"  \
        header \
            -name "X-Header1" -value "custom header1" \
        header \
            -name "X-Header2" -value "custom header2" \
        header \
            -name "Disposition-Notification-To" -value "${FROM}"
