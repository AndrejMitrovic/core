package keeper

import (
	"context"
	"strconv"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/terra-money/core/x/token/types"
	core "github.com/terra-money/core/types"
)

func GetMicroKRW (amountStr string, k msgServer, ctx sdk.Context) (sdk.Dec, error) {
	amount, err := strconv.ParseUint(amountStr, 10, 64);
	if err != nil {
		return sdk.Dec{}, sdkerrors.Wrap(types.ErrInvalidAmount, "Amount must be a non-negative number")
	}

	lunaAmount := sdk.NewInt64DecCoin(core.MicroLunaDenom, int64(amount))
	exRate, err := k.OracleKeeper.GetLunaExchangeRate(ctx, core.MicroKRWDenom)
	if err != nil {
		// use an assumed amount if the Oracle isn't running (just for demonstration purposes)
		exRate = sdk.NewDec(16944000)

		// .. but normally we'd return an error
		// return nil, sdkerrors.Wrap(market.ErrNoEffectivePrice, "No exchange rate for KRW <-> Luna")
	}

	ukrwAmount := lunaAmount.Amount.Mul(exRate)
	if ukrwAmount.LTE(sdk.ZeroDec()) {
		return sdk.Dec{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Value is less than zero")
	}

	return ukrwAmount, nil
}

func (k msgServer) CreateCoins(goCtx context.Context, msg *types.MsgCreateCoins) (*types.MsgCreateCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetCoins(
		ctx,
		msg.User,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	amount, error := GetMicroKRW(msg.Amount, k, ctx)
	if error != nil {
		return nil, error
	}

	var coins = types.Coins{
		Creator: msg.Creator,
		User:    msg.User,
		Amount:  amount.String(),
	}

	k.SetCoins(
		ctx,
		coins,
	)
	return &types.MsgCreateCoinsResponse{}, nil
}

func (k msgServer) UpdateCoins(goCtx context.Context, msg *types.MsgUpdateCoins) (*types.MsgUpdateCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCoins(
		ctx,
		msg.User,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	newAmount, error := GetMicroKRW(msg.Amount, k, ctx)

	if error != nil {
		return nil, error
	}

	// panics on invalid input as validation should have already been done
	getValidAmount := func (input string) sdk.Dec {
		if amount, err := sdk.NewDecFromStr(input); err == nil {
			return amount
		} else {
			panic(fmt.Sprintf("Value is not stored as a decimal: %v", msg.Amount))
		}
	}

	// add previous coins
	prevAmount := getValidAmount(valFound.Amount)
	totalAmount := newAmount.Add(prevAmount)

	var coins = types.Coins{
		Creator: msg.Creator,
		User:    msg.User,
		Amount:  totalAmount.String(),
	}

	k.SetCoins(ctx, coins)

	return &types.MsgUpdateCoinsResponse{}, nil
}

func (k msgServer) DeleteCoins(goCtx context.Context, msg *types.MsgDeleteCoins) (*types.MsgDeleteCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCoins(
		ctx,
		msg.User,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCoins(
		ctx,
		msg.User,
	)

	return &types.MsgDeleteCoinsResponse{}, nil
}
