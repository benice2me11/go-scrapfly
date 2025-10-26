#!/bin/bash

API_KEY=$1

if [ -z "$API_KEY" ]; then
    echo "Usage: $0 <API_KEY>"
    exit 1
fi

run_example() {
    local example=$1
    echo "--------------------------------"
    echo "Running example: $example"
    echo "--------------------------------"
    go run examples/main.go $example $API_KEY
}

get_examples() {
    local examples=$(go run examples/main.go|grep -|cut -d' ' -f3 )
    echo $examples
}

# Run all examples
for example in $(get_examples); do
    run_example $example
done