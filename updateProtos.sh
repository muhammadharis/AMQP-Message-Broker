#!/bin/bash

for f in $(find ./protos -name '*.proto');
do
	echo "Processing $f"
    protoc --go_out=plugins=grpc:. $f
done