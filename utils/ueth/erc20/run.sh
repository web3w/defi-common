#!/bin/bash

# compile
solc token.sol --abi
# generate binding
abigen --abi ./token_sol_Token.abi --pkg erc20 > token.go