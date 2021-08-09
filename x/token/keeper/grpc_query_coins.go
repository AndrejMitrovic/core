package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/terra-money/core/x/token/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CoinsAll(c context.Context, req *types.QueryAllCoinsRequest) (*types.QueryAllCoinsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinss []*types.Coins
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	coinsStore := prefix.NewStore(store, types.KeyPrefix(types.CoinsKeyPrefix))

	pageRes, err := query.Paginate(coinsStore, req.Pagination, func(key []byte, value []byte) error {
		var coins types.Coins
		if err := k.cdc.Unmarshal(value, &coins); err != nil {
			return err
		}

		coinss = append(coinss, &coins)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoinsResponse{Coins: coinss, Pagination: pageRes}, nil
}

func (k Keeper) Coins(c context.Context, req *types.QueryGetCoinsRequest) (*types.QueryGetCoinsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCoins(
		ctx,
		req.User,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetCoinsResponse{Coins: &val}, nil
}
