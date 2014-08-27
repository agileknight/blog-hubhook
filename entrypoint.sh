#!/bin/bash

set -e

[ $(cat /config.ini | grep apikeys) ] || \
echo "
[apikeys]
key = ${API_KEY}
" >> /config.ini

exec ./hub-listener -listen 0.0.0.0:80 -config-file /config.ini "$@"
