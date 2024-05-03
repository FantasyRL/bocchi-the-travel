#!/bin/bash

(go run ./api) &
sleep 2s
printf "\n"

(go run ./rpc/user) &
sleep 2s
printf "\n"

(go run ./rpc/party) &

wait