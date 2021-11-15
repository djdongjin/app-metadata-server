#!/usr/bin/env bash

curl http://localhost:8080/persist \
    --include \
    --header "Content-Type: text/plain" \
    --request "POST" \
    --data "$(cat $1)"