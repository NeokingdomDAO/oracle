# Oracle

The oracle bridges data from Odoo to the contracs in EVMOS.

## Development

- Clone the repository
- Init the submodule: `git submodule init`
- Update: `git submodule update`

## Generate Contract Modules

Below the instructions to generate the contract interactions modules starting from the contract ABIs.

### Requirements

- [`jq`](https://stedolan.github.io/jq/)
- [`abigen`](https://geth.ethereum.org/docs/dapp/native-bindings)

### Steps

Compile the contracts:

```
cd contracts
git checkout v1.0.0
pnpm install
pnpm compile
```

Run the generator

```
# Example with contract ResolutionManager. Replace with required contract.

cat contracts/artifacts/contracts/ResolutionManager/ResolutionManager.sol/ResolutionManager.json | jq ".abi" > ResolutionManager.json

abigen --abi ResolutionManager.json --pkg eth --type ResolutionManager --out ResolutionManager.go

rm ResolutionManager.json
```

Or use the handy `./contracts.sh`.
