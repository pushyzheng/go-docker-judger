#!/usr/bin/env bash
cd code/$1
javac Main.java

if [ $? -ne 0 ]
then
    echo "COMPILE_ERROR"
    exit 1
fi

time java Main < $2 > result.txt
if [ $? -ne 0 ]
then
    echo "RUNTIME_ERROR"
    exit 1
else
    echo "SUCCEED"
fi
