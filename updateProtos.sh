#!/bin/bash

for f in ./api/*.proto
do
	echo "Processing $f"
    protoc --go_out=paths=source_relative:. $f
done