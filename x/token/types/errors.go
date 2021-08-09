package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/token module sentinel errors
var (
	ErrInvalidAmount = sdkerrors.Register(ModuleName, 1100, "Invalid amount")
	// this line is used by starport scaffolding # ibc/errors
)
