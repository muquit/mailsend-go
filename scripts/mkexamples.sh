#!/bin/bash
# generate examples.txt file which gets embedded in binary by packr
# muquit@muquit.com Nov-02-2018 
PANDOC=`which pandoc`
FROM="docs/examples.md"
TO="files/examples.txt"
MYDIR=$(dirname $0)
ME=$(basename $0)
DATE=$(date)
INPUT="./files/examples.txt"
RM="/bin/rm -f"
MV="/bin/mv -f"

CreateExamplesTxt() {
    if [[ -f $PANDOC ]]; then
        echo "- Generating $FROM ==> $TO"
        ${PANDOC} -f markdown -t plain $FROM > $TO
    else
        echo "ERROR: did not find pandoc, will generate examples file"
        exit 1
    fi
}

#########################################################################
# Just generate examples.go from files/examples.txt as string literals.
# I tried packr, packr2 etc, strange bugs introduced in packr2, binary 
# became 11MB and crashes if run with abs path, somehow it pulled too 
# much garbage in the binary.
#
# Also it gets difficult with vendoring. When used glide it can't find the
# packr call. When used dep, it pulled tons of garbage in vendor directory.
# WTH?
# 
# I don't have time to deal with al the bs.
# 
# muquit@muquit.com Jan-20-2019 
#########################################################################

CreateExamplesGo() 
{
    echo "// DO NOT MODIFY - Automatically generated from ${INPUT}"
    echo ""
    echo 'package main


import "fmt"

const (
    examples = `
'
while IFS= read -r var
do
    echo "$var"
done < "${INPUT}"

echo '`
)

// Print Examples ...
func PrintExamples() {
    fmt.Println(examples)
}
'
}

CreateExamplesTxt

OUT=$(mktemp -t "mailsend-go")
TO="./examples.go"
CreateExamplesGo > ${OUT}
cmp -s ${OUT} ${TO}
RC=$?
if [[ $RC = 1 ]]; then
    echo " - Replace examples.go..."
    ${MV} ${OUT} ${TO}
else
    echo " - No need to replace ${TO}"
fi
${RM} ${OUT}
exit 0

