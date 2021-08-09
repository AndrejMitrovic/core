package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/terra-money/core/x/token/types"
)

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

	var coins = types.Coins{
		Creator: msg.Creator,
		User:    msg.User,
		Amount:  msg.Amount,
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

	var coins = types.Coins{
		Creator: msg.Creator,
		User:    msg.User,
		Amount:  msg.Amount,
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
