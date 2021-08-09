## How to run

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
