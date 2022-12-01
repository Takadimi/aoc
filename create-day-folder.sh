#!/bin/bash

DAY=$1

if [[ -z $DAY ]]; then
    echo "day required (e.g. ./create-day-folder.sh 2)"
    exit 1
fi



mkdir day-$DAY
(
    cd day-$DAY &&
    touch main.go sample.txt input.txt &&
    cat << EOF > main.go
package main

func main() {
}
EOF
)
