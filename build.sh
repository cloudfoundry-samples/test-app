#!/usr/bin/env bash

echo "Compiling for linux..."
GOOS=linux GOARCH=amd64 go build .

echo "Constructing Dockerimage"
docker build -t="jbayer/lattice-app" .
docker push jbayer/lattice-app

echo "Cleaning up..."
rm lattice-app
