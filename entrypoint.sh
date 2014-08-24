#!/bin/bash

set -e

echo "
[apikeys]
key = ${API_KEY}
" >> /conf.ini

exec ./hub-listener -listen 0.0.0.0:80 -config-file /conf.ini "$@"
