## What this contains

This adds a very rudimentary token "wallet" module. The user signs a transaction
containing the amount of tokens they want to create / update in microLuna
denomination, and then the module converts it to microKRW via the Oracle module.

I did not add any other useful functionality such as sending tokens from one
account to another, as I've ran out of time due to:
- Researching & reading Terra / Cosmos SDK / Starport documentation
- Refreshing my memory about how Golang works
- Trying to tame Starport to make it do what I wanted. :)
- Resolving version differences, due to Starport master targeting an older
  Cosmos SDK.

I would have loved adding more features (+ unittests), but I needed more time.

## How to run

I've used Go `v1.16.6` on linux.

### Starting terra-core:

```
make install
terrad init --chain-id=testnet usernode
terrad keys add user
terrad add-genesis-account $(terrad keys show user -a) 100000000uluna,1000usd
terrad gentx user 10000000uluna --chain-id=testnet
terrad collect-gentxs
terrad start
```

Testing token functionality:

```
# Create some tokens
terrad tx token create-coins 1000 --from=user --chain-id=testnet --gas=auto

# Add some tokens
terrad tx token update-coins 2500 --from=user --chain-id=testnet --gas=auto

# Should fail validation, negative numbers disallowed
terrad tx token update-coins "\-500" --from=user --chain-id=testnet --gas=auto
```

```
# This should list the stored tokens
$ terrad query token list-coins
```
