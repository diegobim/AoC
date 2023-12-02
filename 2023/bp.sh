#!/usr/bin/env bash

# this generates the boilerplate for a given day

DAY="$1"

if [ -z $DAY ]; then 
	echo "Usage: ./boilerplate.sh <day=[1-25]>"
	exit 1 
fi

if [ -d ./$DAY ]; then 
	echo "Directory already exists"
	exit
fi

cp ./0 ./"$DAY" -r