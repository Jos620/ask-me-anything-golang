#!/bin/bash

if ! command -v docker-compose &> /dev/null; then
    echo "docker-compose could not be found. Please install it first."
    exit 1
fi

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo ".env file not found!"
    exit 1
fi

if ! [ -f ./configs/docker-compose.yaml ]; then
    echo "Docker composer configuration not found!"
    exit 1
fi

docker-compose -f ./configs/docker-compose.yaml up -d
