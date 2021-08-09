package keeper

import (
	"github.com/terra-money/core/x/token/types"
)

var _ types.QueryServer = Keeper{}
