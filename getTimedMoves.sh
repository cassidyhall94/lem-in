#!/bin/bash

for m in $(ls maps | grep -v bad | grep example); do
    start=`date +%s`
    moves=$(go run . maps/$m | wc -l)
    end=`date +%s`

    runtime=$((end-start))

    echo "For farm maps/$m we got $moves total in $runtime seconds"
done