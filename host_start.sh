#!/bin/bash

set -x

cd halbarad 2>/dev/null
mkdir ignored 2>/dev/null 

set -ex

python3 -m http.server 9000 -d web > ignored/server.log 2>&1 &

set -e

echo "Done"