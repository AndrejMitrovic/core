package token

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/core/x/token/keeper"
	"github.com/terra-money/core/x/token/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the coins
	for _, elem := range genState.CoinsList {
		k.SetCoins(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all coins
	coinsList := k.GetAllCoins(ctx)
	for _, elem := range coinsList {
		elem := elem
		genesis.CoinsList = append(genesis.CoinsList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
