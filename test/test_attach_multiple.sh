#!/bin/bash

########################################################################
# Test STMP info
# Part of mailsend-go muquit@muquit.com 
########################################################################

readonly DIR=$(dirname $0)
source ${DIR}/env.txt


$MAILSEND -smtp $SMTP_SERVER -port $TLS_PORT \
    -debug \
    -sub "Mail with a HTML message and a PDF attachment" \
    -t "${TO}" \
    -f "${FROM}" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -msg '
<html>
    <header>
        <title>This is a Test</title>
    </header>
    <body>
        <h1>hello, world!</h1>
        This is the mail body
        <ul>
        <li> PDF file attached
        <li> Cat embedded
        <li> flower attached
        </ul>
        But it is up to the mail reader how they will be displayed!
        <hr>
        The End
    </body>
</html>
' \
    attach -file "$DIR/file.pdf" \
    attach -file "$DIR/cat.jpg" -inline \
    attach -file "$DIR/flower.jpg" -name "FLOWER.JPG"
    
