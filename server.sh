#!/bin/bash

set -x

cd /home/wesley_oates/projects/halbarad 
mkdir ignored

python3 -m http.server 9000 > ignored/server.log 2>&1 &