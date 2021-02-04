#!/usr/bin/env bash
NAME="user_center"

mkdir -p output/bin output/config
cp -r config/ output/config/

go build -o output/bin/${NAME}
chmod +x output/bin/${NAME}
chmod +x output/start.sh