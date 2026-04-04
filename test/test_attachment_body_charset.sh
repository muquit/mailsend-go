#!/bin/bash
########################################################################
# Test attachment MIME type is separate from body charset
########################################################################
readonly DIR=$(dirname $0)
source ${DIR}/env.txt

/bin/rm -f /tmp/test.txt
echo "Test data" > /tmp/test.txt

$MAILSEND -smtp $SMTP_SERVER -port $TLS_PORT \
    -debug \
    -sub "Body charset vs attachment MIME" \
    -t "${TO}" \
    -f "${FROM}" \
    -cs "windows-1250" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -msg "Body with windows-1250" \
    attach -file /tmp/test.txt -mime-type "text/plain"

/bin/rm /tmp/test.txt
echo "Verify body has charset=windows-1250 but attachment has its own MIME"
