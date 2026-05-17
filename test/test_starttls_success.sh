#!/bin/bash

########################################################################
# Test successful STARTTLS delivery against Mailpit. 
# mailsend-go v1.0.12+
# Mailpit advertises STARTTLS when started with TLS cert/key:
#   mailpit --smtp-tls-cert cert.pem --smtp-tls-key key.pem
# mailsend-go. Pass -verifyCert if using a trusted cert.
# Expected result: mailsend-go exits 0 and the message appears in the
# Mailpit web UI (default: http://localhost:8025).
# Part of mailsend-go muquit@muquit.com
# Start mailpit like:
# mailpit --smtp-tls-cert example.org.pem --smtp-tls-key example.org-key.pem --smtp-auth-file auth.txt
#  use mkcert https://github.com/FiloSottile/mkcert to create cert
#  cat auth.txt
#  test:test
########################################################################

readonly DIR=$(dirname "$0")

MAILSEND="${MAILSEND:-./mailsend-go}"
MAILPIT_HOST="${MAILPIT_HOST:-localhost}"
MAILPIT_PORT="${MAILPIT_PORT:-1025}"

echo "=== STARTTLS success test (Mailpit) ==="
echo "Server : ${MAILPIT_HOST}:${MAILPIT_PORT}"
echo "Expecting mailsend-go to SUCCEED — Mailpit advertises STARTTLS."
echo ""

output=$("$MAILSEND" \
    -smtp  "$MAILPIT_HOST" \
    -port  "$MAILPIT_PORT" \
    -sub   "STARTTLS success test" \
    -t     "test@example.com" \
    -f     "sender@example.com" \
    auth -user test -pass test \
    body -msg "This message was delivered via STARTTLS." 2>&1)

rc=$?
echo "$output"
echo ""

if [[ $rc -eq 0 ]]; then
    echo "PASS: mailsend-go delivered via STARTTLS (exit $rc)."
    exit 0
else
    echo "FAIL: mailsend-go failed to deliver via STARTTLS (exit $rc)."
    exit 1
fi
