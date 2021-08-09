package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// OracleKeeper defines expected oracle keeper
type OracleKeeper interface {
    GetLunaExchangeRate(ctx sdk.Context, denom string) (price sdk.Dec, err error)
    GetTobinTax(ctx sdk.Context, denom string) (tobinTax sdk.Dec, err error)

    // only used for simulation
    IterateLunaExchangeRates(ctx sdk.Context, handler func(denom string, exchangeRate sdk.Dec) (stop bool))
    SetLunaExchangeRate(ctx sdk.Context, denom string, exchangeRate sdk.Dec)
    SetTobinTax(ctx sdk.Context, denom string, tobinTax sdk.Dec)
}
