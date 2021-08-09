package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/terra-money/core/x/token/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCoinsMsgServerCreate(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	srv := NewMsgServerImpl(*keeper)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateCoins{Creator: creator,
			User: strconv.Itoa(i),
		}
		_, err := srv.CreateCoins(wctx, expected)
		require.NoError(t, err)
		rst, found := keeper.GetCoins(ctx,
			expected.User,
		)
		require.True(t, found)
		assert.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestCoinsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateCoins
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateCoins{Creator: creator,
				User: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateCoins{Creator: "B",
				User: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateCoins{Creator: creator,
				User: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			keeper, ctx := setupKeeper(t)
			srv := NewMsgServerImpl(*keeper)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateCoins{Creator: creator,
				User: strconv.Itoa(0),
			}
			_, err := srv.CreateCoins(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateCoins(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := keeper.GetCoins(ctx,
					expected.User,
				)
				require.True(t, found)
				assert.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestCoinsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteCoins
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteCoins{Creator: creator,
				User: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteCoins{Creator: "B",
				User: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteCoins{Creator: creator,
				User: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			keeper, ctx := setupKeeper(t)
			srv := NewMsgServerImpl(*keeper)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateCoins(wctx, &types.MsgCreateCoins{Creator: creator,
				User: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteCoins(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := keeper.GetCoins(ctx,
					tc.request.User,
				)
				require.False(t, found)
			}
		})
	}
}
