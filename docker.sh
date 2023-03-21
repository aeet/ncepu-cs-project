#!bin/bash

docker build -t devcui-nceps-cs-project -f . 
docker tag devcui-nceps-cs-project devcui-nceps-cs-project:v1
docker push devcui-nceps-cs-project