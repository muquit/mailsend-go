#!/bin/bash
########################################################################
# Test default charset (UTF-8) with HTML
########################################################################
readonly DIR=$(dirname $0)
source ${DIR}/env.txt
$MAILSEND -smtp $SMTP_SERVER -port $TLS_PORT \
    -debug \
    -sub "HTML default UTF-8" \
    -t "${TO}" \
    -f "${FROM}" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -msg '<html><body><p>Unicode: 你好 مرحبا 🎉</p></body></html>'

echo "Verify headers show: Content-Type: text/html; charset=UTF-8"