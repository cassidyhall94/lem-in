#!/bin/bash

for m in $(ls maps | grep -v bad | grep example); do
    echo "-------maps/$m-------"
    go run . maps/$m
    echo "--------------------------------"
done