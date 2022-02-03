package cmd

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc20/config"
	"github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20/erc20mgr/types"
	"github.com/spf13/cobra"
)

func fabricCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "fabric",
		Short:                      "Fabric subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		clientIDFabricCmd(ctx),
		idFabricCmd(ctx),
	)

	return cmd
}

func clientIDFabricCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clientId",
		Short: "Get client id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientID, err := clientIDFromWallet(homePath, ctx.Config.WalletLabel)
			if err != nil {
				return err
			}
			fmt.Println(clientID)

			return nil
		},
	}

	return cmd
}

func idFabricCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "id",
		Short: "Get id in contract module",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientID, err := clientIDFromWallet(homePath, ctx.Config.WalletLabel)
			if err != nil {
				return err
			}

			id, err := types.NewID(types.MSPID(ctx.Config.WalletLabel), types.ClientID(clientID))
			if err != nil {
				return err
			}
			fmt.Println(id)

			return nil
		},
	}

	return cmd
}
