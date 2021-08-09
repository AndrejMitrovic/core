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
		Use:   "create-coins [user] [amount]",
		Short: "Create a new coins",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get indexes
			indexUser, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			// Get value arguments
			argsAmount, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCoins(
				clientCtx.GetFromAddress().String(),
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
		Use:   "update-coins [user] [amount]",
		Short: "Update a coins",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get indexes
			indexUser, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			// Get value arguments
			argsAmount, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateCoins(
				clientCtx.GetFromAddress().String(),
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

func CmdDeleteCoins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-coins [user]",
		Short: "Delete a coins",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			indexUser, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteCoins(
				clientCtx.GetFromAddress().String(),
				indexUser,
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
