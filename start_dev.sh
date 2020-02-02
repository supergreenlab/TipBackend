#!/bin/bash

docker build -t tipserver-dev . -f Dockerfile.dev
docker run  --name=tipserver --network=supergreencloud_back-tier -p 8080:8080 --rm -it -v $(pwd)/config:/etc/tipserver -v $(pwd):/app tipserver-dev
docker rmi tipserver-dev
