#!/bin/bash

set -ex

pid=$(ps ax | grep "python3 -m http.server 9000 -d web" | grep -v grep | sed 's/|/ /' | awk '{print $1}')
kill $pid

set -e

echo "Done"