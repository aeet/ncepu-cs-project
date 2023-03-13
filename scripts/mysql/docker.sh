#!/bin/bash

podman run \
    --restart=always \
    --name mysql \
    -v /opt/mysql/conf/my.cnf:/etc/mysql/my.cnf:rw \
    -v /opt/mysql/data:/var/lib/mysql \
    -v /opt/mysql/logs:/var/log/mysql \
    -v /etc/localtime:/etc/localtime:ro \
    -p 3306:3306 \
    -e TZ=Asia/Shanghai \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -d mysql \
    --character-set-server=utf8mb4 \
    --collation-server=utf8mb4_general_ci
