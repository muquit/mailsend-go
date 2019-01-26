#!/bin/bash

########################################################################
# Test STMP info
# Part of mailsend-go muquit@muquit.com 
########################################################################

readonly DIR=$(dirname $0)
source ${DIR}/env.txt

$MAILSEND -smtp $SMTP_SERVER -port $TLS_PORT \
    -debug \
    -sub "Mail with a HTML message and an embedded image" \
    -t "${TO}" \
    -f "${FROM}" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -msg '
<html>
    <header>
        <title>This is a Test</title>
    </header>
    <body>
    <h2>hello, world!</h2>
    A flower is embedded with Content-Disposition to inline. The mail reader
    should show it on the page
        <hr>
        The End
    </body>
</html>
' \
    attach -file "$DIR/flower.jpg" -inline
    
    
