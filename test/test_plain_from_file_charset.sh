#!/bin/bash
########################################################################
# Test plain text from file with charset
########################################################################
readonly DIR=$(dirname $0)
source ${DIR}/env.txt

cat > /tmp/test.txt << 'EOF'
This is plain text.
Line 2 with special chars: é à ü
Line 3
EOF

$MAILSEND -smtp $SMTP_SERVER -port $TLS_PORT \
    -debug \
    -sub "Text file with ISO-8859-15" \
    -t "${TO}" \
    -f "${FROM}" \
    -cs "ISO-8859-15" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -file /tmp/test.txt

rm /tmp/test.txt
echo "Verify: Content-Type: text/plain; charset=ISO-8859-15"