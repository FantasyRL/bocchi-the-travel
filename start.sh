#!/bin/bash

(go run ./api) &
sleep 2s
printf "\n"

(go run ./rpc/user) &
sleep 2s
printf "\n"

(go run ./rpc/party) &
sleep 2s
printf "\n"

(go run ./rpc/itinerary) &
sleep 2s
printf "\n"

(go run ./rpc/interaction) &
sleep 2s
printf "\n"

(go run ./rpc/trust) &
wait