#!/bin/bash

########################################################################
# Test STMP info
# Part of mailsend-go muquit@muquit.com 
########################################################################
#
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
        <h1>Hello</h1>
        Unordered list:
        <ul>
            <li>one</li>
            <li>two</li>
            <li>three</li>
        </ul>
        To convert in pdf:
        <pre>
pandoc file.html --pdf-engine=pdflatex -o file.pdf
        </pre>
        <hr>
        The End
    </body>
</html>
' \
    attach -file "$DIR/file.pdf"
    
