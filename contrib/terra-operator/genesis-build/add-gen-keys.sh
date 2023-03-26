#!/bin/sh

KEYRING="${KEYRING:-test}"

# add keys, add balances
for i in $(seq 0 3); do
    key=$(jq ".keys[$i] | tostring" /keys.json )
    keyname=$(echo $key | jq -r 'fromjson | ."keyring-keyname"')
    mnemonic=$(echo $key | jq -r 'fromjson | .mnemonic')
    # Add new account
    echo $mnemonic | terrad keys add $keyname --keyring-backend $KEYRING --recover --home ~/.terra
    # Add initial balances
    terrad add-genesis-account $keyname "1000000000000uluna" --keyring-backend $KEYRING --home ~/.terra
done