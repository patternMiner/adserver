#!/bin/bash

#
# Created by jbisa on 8/20/17.
#

/usr/local/go/bin/go run /Users/jbisa/go/src/github.com/patternMiner/adserver/main.go \
    --cert_path=$HOME/.ssh/cert.pem --key_path=$HOME/.ssh/key.pem
