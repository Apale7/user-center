#!/usr/bin/env bash
NAME="user_center"
#编译pb
# for x in **/*.proto; 
# do protoc --go_out=plugins=grpc,paths=source_relative:. $x; done

mkdir -p output/bin output/config
cp -r config/ output/config/

go build -o output/bin/${NAME}
chmod +x output/bin/${NAME}
chmod +x output/start.sh