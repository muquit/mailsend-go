#!/bin/bash

########################################################################
# Test STMP info
# Part of mailsend-go muquit@muquit.com 
########################################################################

readonly DIR=$(dirname $0)

LOGFILE="/tmp/test$$.log"
/bin/rm -f ${LOGFILE}

$MAILSEND \
    -q \
    -sub "Test log file" \
    -t "${TO}" \
    -f "${FROM}" \
    auth -user "${FROM}" -pass "${SMTP_USER_PASS}" \
    body -msg "This is a test" \
    -smtp $SMTP_SERVER -port $TLS_PORT \
    -log ${LOGFILE}
echo "--------------------------logfile--"
cat $LOGFILE
echo "--------------------------logfile--"
/bin/rm -f ${LOGFILE}
