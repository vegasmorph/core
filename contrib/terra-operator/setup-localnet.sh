#!/bin/bash

terrad testnet --keyring-backend=test --v=4 --node-daemon-home=.terra --home=temp --chain-id=$CHAINID
chmod -R 777 mytestnet