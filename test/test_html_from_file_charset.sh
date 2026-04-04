#!/bin/bash
########################################################################
# Test HTML from file with charset
########################################################################
readonly DIR=$(dirname $0)
source ${DIR}/env.txt

/bin/rm -f /tmp/tst.sh
cat > /tmp/test.html << 'EOF'
<html>
<head><meta charset="UTF-8"></head>
<body>
    <h1>File Test</h1>
    <p>Special chars: ñ ü ö</p>
</body>
</html>
EOF

$MAILSEND -smtp $SMTP_SERVER -port $TLS_PORT \
    -debug \
    -sub "HTML file with windows-1252" \
    -t "${TO}" \
    -f "${FROM}" \
    -cs "windows-1252" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -file /tmp/test.html

rm /tmp/test.html
echo "Verify: Content-Type: text/html; charset=windows-1252"
