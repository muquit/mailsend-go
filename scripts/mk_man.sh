#!/bin/bash

# Read version from VERSION file
VERSION=$(cat VERSION | tr -d '\n')

# Get current month and year
MONTH_YEAR=$(date '+%B %Y')

(echo '---'; \
 echo 'title: MAILSEND-GO'; \
 echo 'section: 1'; \
 echo 'header: User Manual'; \
 echo "footer: mailsend-go ${VERSION}"; \
 echo "date: ${MONTH_YEAR}"; \
 echo '---'; \
 echo ''; \
 sed '/^\[\!\[/,/^$/d; /^## Table Of Contents/,/^$/d' README.md) | \
 pandoc -s -t man -o docs/mailsend-go.1

echo "Generated docs/mailsend-go.1 with version ${VERSION} dated ${MONTH_YEAR}"