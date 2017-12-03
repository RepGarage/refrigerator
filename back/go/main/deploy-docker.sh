#!/bin/bash

curl -H "Content-Type: application/json" --data '{"docker_tag": "back-go-main"}' -X POST https://registry.hub.docker.com/u/repgarage/refrigerator/trigger/2064e54f-14c5-4541-a66d-fc3cd0540cfe/