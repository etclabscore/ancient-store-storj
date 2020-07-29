#!/bin/sh

[ ! -d _workspace ] && mkdir -p _workspace
[ ! -d _workspace/go-ethereum ] && git clone https://github.com/etclabscore/core-geth.git _workspace/go-ethereum
cd _workspace/go-ethereum; git fetch origin feat/frozen-remote; git fetch --tags; git checkout origin/feat/frozen-remote; cd ..