#!/bin/bash

########################################################################
# Test STMP info
# Part of mailsend-go muquit@muquit.com 
########################################################################

readonly DIR=$(dirname $0)
source ${DIR}/env.txt


$MAILSEND -info -smtp $SMTP_SERVER -port $TLS_PORT
$MAILSEND -info -smtp $SMTP_SERVER -port $SSL_PORT -ssl
