#!/bin/bash

while true
do
    go build main.go 
    chmod +x main
    ./main 2>/dev/null
    sleep 60
done
