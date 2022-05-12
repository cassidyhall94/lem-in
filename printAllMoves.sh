#!/bin/bash

for m in $(ls maps | grep -v bad | grep example); do
    echo "-------maps/$m-------"
    bin/lem-in maps/$m
    echo "--------------------------------"
done