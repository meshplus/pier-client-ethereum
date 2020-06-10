#!/usr/bin/env sh
set -e

APPCHAIN_NAME=$1

pier --repo=/root/.pier appchain register --name=${APPCHAIN_NAME} --type=ether --validators=/root/.pier/ether/ether.validators --desc="appchain for test" --version=1.9.3
pier --repo=/root/.pier rule deploy --path=/root/.pier/validating.wasm
pier --repo=/root/.pier start