#!/bin/bash

mkdir -p bin

go build -o bin/lem-in . 

chmod +x bin/lem-in

echo "lem-in built, run it with a map like: 'bin/lem-in maps/example00.txt'"