#!/bin/bash

set -e
echo "Building Golang backend API..."
cd api
go build .
cd ..
echo "Checking if the web server is running..."
pid_on_9000=$(lsof -Pi | grep python3 | awk '{print $2}')
cmd_on_9000=$(ps -ef | awk -v pid="$pid_on_9000" '$2==pid { for (i=8; i<=NF; i++) print $i }'  | tr '\n' ' ')
if [[ $cmd_on_9000 == *"http.server"* ]] && [[ $cmd_on_9000 == *"web"* ]] ; then
  echo "Shutting down existing web server..."
  kill $pid_on_9000
  echo "...done"
else
  echo "No existing server."
fi

echo "Starting new server..."
python3 -m http.server 9000 -d web > ignored/server.log 2>&1 &
echo "Done"
