#!/usr/bin/env bash

# this executes the code for a given day and part

DAY="$1"
PART="$2"
TEST="$3"

if [ -z $DAY ]; then 
	echo "Usage: ./run.sh <day=[1-25]> <part=[1-2]> [<test=*>]"
	exit 1 
fi

INPUT_FILE="input.txt"
if [ -n "$TEST" ]; then
		INPUT_FILE="sample.txt"
fi

COMMAND="go run ./$DAY/part$PART/main.go < ./$DAY/$INPUT_FILE"

echo $COMMAND
eval $COMMAND
