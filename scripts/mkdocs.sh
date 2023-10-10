#!/bin/sh
# muquit@muquit.com t-31-2018  
MH="markdown_helper"
RM="/bin/rm -f"
DOC_DIR="./docs"
MAILSEND="./mailsend-go"

if [ -f $MAILSEND ]; then
    echo " - Generate $DOC_DIR/usage.txt"
    ./mailsend-go -h > $DOC_DIR/usage.txt
fi
pushd $DOC_DIR >/dev/null 
echo " - Assembling README.md"
${MH} include --pristine main.md ../README.md
${MH} include --pristine chl.md ../ChangeLog.md
popd >/dev/null
