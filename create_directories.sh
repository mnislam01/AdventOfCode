#!/usr/bin/env bash

for ((index=3; index<=25; index++)); do
  mkdir "day$index"
  echo "Created directory: day$index"
  cd "day$index"
  mkdir "part1"
  echo "../part1"
  mkdir "part2"
  echo "../part2"
  cd ..
done
