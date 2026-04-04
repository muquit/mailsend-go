#!/bin/bash
########################################################################
# Test no charset specified (should default to UTF-8)
########################################################################
readonly DIR=$(dirname $0)
source ${DIR}/env.txt
$MAILSEND -smtp $SMTP_SERVER -port $TLS_PORT \
    -debug \
    -sub "Default charset test" \
    -t "${TO}" \
    -f "${FROM}" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -msg "No charset specified, should be UTF-8"

echo "Verify: Content-Type: text/plain; charset=UTF-8"