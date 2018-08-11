#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o app.o

eval $(minikube docker-env)

docker build -t cotb2018:1.0.0 -f Dockerfile .