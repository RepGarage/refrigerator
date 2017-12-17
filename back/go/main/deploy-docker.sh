#!/bin/bash

curl -H "Content-Type: application/json" --data '{"docker_tag": "back-go"}' -X POST https://registry.hub.docker.com/u/centrypoint/refrigerator/trigger/22d7b61f-d94d-4bfc-97fe-c9ba7035e8e4/