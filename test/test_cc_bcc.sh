#!/bin/bash

readonly DIR=$(dirname $0)
source ${DIR}/env.txt

echo $Cc
echo $BCC

$MAILSEND -debug \
        -subject "Testing Cc and Bcc" \
        -from "${FROM}" \
        -to "$TO" \
        -cc "$Cc" \
        -bcc "$BCc" \
        -smtp "$SMTP_SERVER" \
        -port $TLS_PORT \
        auth -user "$SMTP_USER" -pass "$SMTP_USER_PASS" \
        body -msg "
        <h2>hello, world!</h2>
        This mail to Cc'd to one address abd Bcc'd  to another address.
        You should only see the Cc'd address
        " \
        -mime-type "text/html" 
