#!/bin/bash

function mongo {
    docker-compose -f docker-compose-dev.yml up -d mongo.local   
    sleep 5
    
    containerId=$(docker ps | grep mongo |awk  '{print $1}')
    docker exec -it $containerId mongo auth --eval "db.users.remove({})"
    docker exec -it $containerId mongo auth --eval "db.users.insert({'username':'example', 'password':'example', 'id':1})"
}

function dep {
    docker-compose -f docker-compose-dev.yml up dep
}

function build {
    docker build -t go-tools .
}

function serve {
    build && dep && mongo && run
}

function run {
    docker-compose -f docker-compose-dev.yml up app
}

"$@"