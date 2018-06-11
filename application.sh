#!/bin/bash

function mongo {
    docker-compose -f docker-compose-dev.yml up -d mongo.local   
    sleep 5
    docker exec -it app_mongo.local_1 mongo auth --eval "db.users.remove({})"
    docker exec -it app_mongo.local_1 mongo auth --eval "db.users.insert({'username':'example', 'password':'example'})"
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