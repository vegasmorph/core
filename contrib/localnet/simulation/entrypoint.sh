#!/bin/sh

export HOME=${HOME:-~/.terra}
export TESTNET_FOLDER=${TESTNET_FOLDER:-build}
export SIMULATION_FOLDER=${SIMULATION_FOLDER:-contrib/localnet/simulation}
export KEYRING_BACKEND=test
export CHAIN_ID=${CHAIN_ID:-localterra}

# initialize keys
for i in $(seq 0 3); do
    key=$(jq ".keys[$i] | tostring" $SIMULATION_FOLDER/keys.json )
    keyname=$(echo $key | jq -r 'fromjson | ."keyring-keyname"')
    mnemonic=$(echo $key | jq -r 'fromjson | .mnemonic')
    # Add new account
    echo $mnemonic | terrad keys add $keyname --keyring-backend $KEYRING_BACKEND --home $HOME --recover
done

# tx_send
sh $SIMULATION_FOLDER/tx_send.sh

echo "DONE TX SEND SIMULATION (1/5)"

# delegate
sh $SIMULATION_FOLDER/delegate.sh

echo "DONE DELEGATION SIMULATION (2/5)"

# create-validator
sh $SIMULATION_FOLDER/create-validator.sh

echo "DONE CREATE VALIDATOR SIMULATION (3/5)"

# contracts
sh $SIMULATION_FOLDER/contract.sh

echo "DONE CONTRACT SIMULATION (4/5)"