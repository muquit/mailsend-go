#!/bin/bash

########################################################################
# Test STARTTLS downgrade protection against MailHog.
# mailsend-go v1.0.12+
# MailHog does not advertise STARTTLS. With RequireSTARTTLS enforced,
# mailsend-go must refuse the connection instead of falling back to
# plaintext and leaking credentials.
# Expected result: mailsend-go exits non-zero with a STARTTLS error.
# Part of mailsend-go muquit@muquit.com
########################################################################

readonly DIR=$(dirname "$0")

MAILSEND="${MAILSEND:-./mailsend-go}"
MAILHOG_HOST="${MAILHOG_HOST:-localhost}"
MAILHOG_PORT="${MAILHOG_PORT:-1025}"

echo "=== STARTTLS downgrade-protection test (MailHog) ==="
echo "Server : ${MAILHOG_HOST}:${MAILHOG_PORT}"
echo "Expecting mailsend-go to FAIL — MailHog does not advertise STARTTLS."
echo ""

output=$("$MAILSEND" \
    -smtp  "$MAILHOG_HOST" \
    -port  "$MAILHOG_PORT" \
    -sub   "STARTTLS protection test" \
    -t     "test@example.com" \
    -f     "sender@example.com" \
    auth -user test -pass test \
    body -msg "This message should never be delivered." 2>&1)

rc=$?
echo "$output"
echo ""

if [[ $rc -ne 0 ]]; then
    echo "PASS: mailsend-go refused connection (exit $rc) — STARTTLS protection is working."
    exit 0
else
    echo "FAIL: mailsend-go delivered without STARTTLS (exit $rc) — protection is NOT working."
    exit 1
fi
