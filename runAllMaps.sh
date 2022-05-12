#!/bin/bash

set +e

for m in $(ls maps); do
    echo "-------maps/$m-------"
    bin/lem-in maps/$m 2>&1
    echo "--------------------------------"
done