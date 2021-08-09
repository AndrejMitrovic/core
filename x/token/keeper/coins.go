package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/core/x/token/types"
)

// SetCoins set a specific coins in the store from its index
func (k Keeper) SetCoins(ctx sdk.Context, coins types.Coins) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKeyPrefix))
	b := k.cdc.MustMarshal(&coins)
	store.Set(types.CoinsKey(
		coins.User,
	), b)
}

// GetCoins returns a coins from its index
func (k Keeper) GetCoins(
	ctx sdk.Context,
	user string,

) (val types.Coins, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKeyPrefix))

	b := store.Get(types.CoinsKey(
		user,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCoins removes a coins from the store
func (k Keeper) RemoveCoins(
	ctx sdk.Context,
	user string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKeyPrefix))
	store.Delete(types.CoinsKey(
		user,
	))
}

// GetAllCoins returns all coins
func (k Keeper) GetAllCoins(ctx sdk.Context) (list []types.Coins) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Coins
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
