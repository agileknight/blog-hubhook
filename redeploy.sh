#!/bin/bash

cd /home/testuser/ts
sudo fig run --rm addrapp rake db:migrate
sudo fig up -d