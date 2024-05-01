#!/bin/bash

# set -x # this is verbose
# set -e # this is strict

which go
which something

if ! [ -f /usr/bin/baus_server ]
then
    go build -o /usr/bin/baus_server ./server/main.go
else
    echo "/usr/bin/baus_server already exists"
fi


if ! [[ -d /etc/baus/data ]]
then
   mkdir -p /etc/baus/data/
   cp config.example.json /etc/baus/config.json
else
    echo "baus directory already exists"
fi



       
sed -i 's/\/tmp\/kjv.json/\/etc/\/baus\/data\/g' /etc/baus/config.json


