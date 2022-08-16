#!/bin/bash

for c in $(ls contracts/artifacts/contracts | grep "[A-Z].*");
do
    cat contracts/artifacts/contracts/$c/$c.sol/$c.json | jq ".abi" > $c.json
    abigen --abi $c.json --pkg eth --type $c --out eth/$c.go
    rm $c.json
done