package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/terra-money/core/x/token/types"
)

func CmdListCoins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-coins",
		Short: "list all coins",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllCoinsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CoinsAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowCoins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-coins [user]",
		Short: "shows a coins",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argsUser, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetCoinsRequest{
				User: argsUser,
			}

			res, err := queryClient.Coins(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
