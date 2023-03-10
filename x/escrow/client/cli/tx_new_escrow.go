package cli

import (
	"strconv"

	"escrow/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdNewEscrow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-escrow [strategy] [instigator] [instigator-wager] [rider] [rider-wager]",
		Short: "Broadcast message newEscrow",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStrategy := args[0]
			argInstigator := args[1]
			argInstigatorWager := args[2]
			argRider := args[3]
			argRiderWager := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgNewEscrow(
				clientCtx.GetFromAddress().String(),
				argStrategy,
				argInstigator,
				argInstigatorWager,
				argRider,
				argRiderWager,
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
