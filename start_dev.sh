#!/bin/bash

docker build -t tipbackend-dev . -f Dockerfile.dev
docker run  --name=tipbackend --network=supergreencloud_back-tier -p 8080:8080 --rm -it -v $(pwd)/config:/etc/tipbackend -v $(pwd):/app tipbackend-dev
docker rmi tipbackend-dev
