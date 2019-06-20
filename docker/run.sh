#!/usr/bin/env bash
javac $1

if [[ $? -ne 0 ]]
then 
    echo "CE"
    exit 1
fi

cd code

java Main < $2 > result.txt
if [[ $? -ne 0 ]]
then
    echo "RE"
    exit 1
fi

cmp ../answer.txt result.txt
if [[ $? -eq 0 ]]
then
    echo "AC"
else
    echo "WA"
fi

cat result.txt