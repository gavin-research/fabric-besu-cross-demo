package cmd

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc20/config"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc20/cross"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc20/cross/atomic"
)

func crossCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "cross",
		Short:                      "Cross subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		cross.GetCreateContractTransaction(ctx),
		cross.QueryTxAuthStateCmd(ctx),
		cross.NewCreateInitiateTxCmd(ctx),
		cross.NewSendInitiateTxCmd(ctx),
		cross.NewSignTxCmd(ctx),
		atomic.GetCoordinatorState(ctx),
	)

	return cmd
}
