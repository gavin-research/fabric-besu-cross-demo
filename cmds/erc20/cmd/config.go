package cmd

import (
	"github.com/spf13/cobra"

	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc20/config"
)

func configCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "manage configuration file",
	}

	cmd.AddCommand(
		config.InitConfigCmd(ctx),
	)

	return cmd
}
