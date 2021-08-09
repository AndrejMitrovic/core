package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/terra-money/core/x/token/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCoins(keeper *Keeper, ctx sdk.Context, n int) []types.Coins {
	items := make([]types.Coins, n)
	for i := range items {
		items[i].User = strconv.Itoa(i)

		keeper.SetCoins(ctx, items[i])
	}
	return items
}

func TestCoinsGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoins(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCoins(ctx,
			item.User,
		)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestCoinsRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoins(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoins(ctx,
			item.User,
		)
		_, found := keeper.GetCoins(ctx,
			item.User,
		)
		assert.False(t, found)
	}
}

func TestCoinsGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCoins(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllCoins(ctx))
}
