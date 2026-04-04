#!/bin/bash
########################################################################
# Test exotic charset
########################################################################
readonly DIR=$(dirname $0)
source ${DIR}/env.txt
$MAILSEND -smtp $SMTP_SERVER -port $TLS_PORT \
    -debug \
    -sub "KOI8-R charset" \
    -t "${TO}" \
    -f "${FROM}" \
    -cs "KOI8-R" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -msg "Testing KOI8-R: Привет"

echo "Verify: Content-Type: text/plain; charset=KOI8-R"