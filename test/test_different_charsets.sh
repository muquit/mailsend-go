#!/bin/bash
########################################################################
# Test mixed: HTML body + plain text attachment with different charsets
########################################################################
readonly DIR=$(dirname $0)
source ${DIR}/env.txt

/bin/rm -f /tmp/plain.txt
echo "Plain attachment content" > /tmp/plain.txt

$MAILSEND -smtp $SMTP_SERVER -port $TLS_PORT \
    -debug \
    -sub "Mixed content types" \
    -t "${TO}" \
    -f "${FROM}" \
    -cs "ISO-8859-1" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -msg '<html><body>HTML body</body></html>' \
    attach -file /tmp/plain.txt

/bin/rm -f /tmp/plain.txt
echo "Verify body is text/html with charset=ISO-8859-1"
