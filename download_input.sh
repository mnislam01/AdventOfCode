#!/usr/bin/env bash

echo $1
echo $2
cd "$1$2"
cd "part1"
ls
curl -x GET "https://adventofcode.com/input/$1/$2/" > input.txt
cp input.txt ../part2
