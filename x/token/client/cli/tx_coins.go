package cli

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/terra-money/core/x/token/types"
)

func CmdCreateCoins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-coins [amount]",
		Short: "Create new coins",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get value arguments
			argsAmount, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			indexUser := clientCtx.GetFromAddress().String()
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCoins(
				indexUser,
				indexUser,
				argsAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateCoins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-coins [amount]",
		Short: "Add more coins",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get value arguments
			argsAmount, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			indexUser := clientCtx.GetFromAddress().String()
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateCoins(
				indexUser,
				indexUser,
				argsAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
