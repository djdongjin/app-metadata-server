#!/usr/bin/env bash

cmd=$1

if [ $cmd == "persist" ]
then
  filename=$2

  curl http://localhost:8080/persist \
    --include \
    --header "Content-Type: text/plain" \
    --request "POST" \
    --data "$(cat $filename)"
elif [ $cmd == "retrieve" ]
then
  query=$2

  curl http://localhost:8080/retrieve \
      --include \
      --header "Content-Type: text/plain" \
      --request "POST" \
      --data "$(echo $query)"
elif [ $cmd == "get" ] || [ $cmd == "delete" ]
then
  title=$2

  curl http://localhost:8080/$cmd/$title
else
  echo "unrecognized subcommand. Select from [persist|retrieve|get|delete]"
fi