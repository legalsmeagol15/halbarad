#!/bin/bash

set -e

# TODO:  use the --cafile or --keyfile or --certfile options to lock down requests
gunicorn --workers 4  --reload --daemon --access-logfile '../ignored/gunicorn.log'  --bind '127.0.0.1:9001' server:app