#!/bin/bash
podman run \
        --restart=always \
        -d \
        -p 6379:6379 \
        -v /opt/redis/redis.conf:/usr/local/etc/redis/redis.conf \
        --privileged=true \
        --name redis redis redis-server /usr/local/etc/redis/redis.conf