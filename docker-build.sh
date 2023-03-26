#!/bin/bash

podman build -t devcui-nceps-cs-project -f ./Dockerfile 
podman tag devcui-nceps-cs-project devcui-nceps-cs-project:v1
podman push devcui-nceps-cs-project