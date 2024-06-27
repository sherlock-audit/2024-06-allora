#!/bin/bash

# Configure predefined mnemonic pharses
BINARY=./rly
CHAIN_DIR=./data
RELAYER_DIR=./relayer

# Ensure rly is installed
if ! [ -x "$(command -v $BINARY)" ]; then
    echo "$BINARY is required to run this script..."
    echo "You can download at https://github.com/cosmos/relayer"
    exit 1
fi

echo "Linking both chains"
$BINARY tx link allora_demo-axelar_demo --home $CHAIN_DIR/$RELAYER_DIR
