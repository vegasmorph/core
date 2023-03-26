#!/bin/bash

terrad testnet --keyring-backend=test --v=$REPLICA --node-daemon-home=.terra --home=temp --chain-id=$CHAINID
chmod -R 777 mytestnet