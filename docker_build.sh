#!/usr/bin/env bash

# docker images list
docker build -t adserver:v1 .

# docker run -d -p 8080:8080 adserver:v1
# docker ps
# docker stop <container-id>
