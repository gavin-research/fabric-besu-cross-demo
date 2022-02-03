package cmd

import (
	"fmt"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc721/config"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc721/erc721"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc721/erc721/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"math/big"
)

func erc721Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "erc721",
		Short: "Operate erc721 token",
	}

	cmd.AddCommand(
		getAddressERC721Cmd(ctx),
		mintERC721Cmd(ctx),
		approveERC721Cmd(ctx),
		getApprovedERC721Cmd(ctx),
		balanceOfERC721Cmd(ctx),
		ownerOfERC721Cmd(ctx),
	)
	return cmd
}

func setupERC721CMD(ctx *config.Context) (erc721.ERC721CMD, error) {
	cmdCfg := ctx.Config
	conn, err := erc721.Connect(cmdCfg.BlockchainHost)
	if err != nil {
		return nil, err
	}

	token, err := contract.NewERC721Token(common.HexToAddress(cmdCfg.ERC721TokenAddress), conn)
	if err != nil {
		return nil, err
	}

	pvtKey, err := cmdCfg.PrivateKey()
	if err != nil {
		return nil, err
	}

	return erc721.NewERC721CMDImpl(conn, cmdCfg.ChainID, pvtKey, token), nil
}

func getAddressERC721Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address",
		Short: "Get Contract address",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(ctx.Config.ERC721TokenAddress)
			return nil
		},
	}

	return cmd
}

func mintERC721Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint",
		Short: "Mint token",
		RunE: func(cmd *cobra.Command, args []string) error {
			erc721CMD, err := setupERC721CMD(ctx)
			if err != nil {
				return err
			}

			receiver, err := cmd.Flags().GetString(FlagAddress)
			if err != nil {
				return err
			}
			tokenID, err := cmd.Flags().GetString(FlagtokenID)
			if err != nil {
				return err
			}
			i, ok := new(big.Int).SetString(tokenID, 10)
			if !ok {
				return errors.New("failed to convert")
			}

			tx, err := erc721CMD.Mint(common.HexToAddress(receiver), i)
			if err != nil {
				return err
			}

			printTx(tx)
			return nil
		},
	}

	cmd.Flags().StringP(FlagAddress, "a", "", "receiver address taken token")
	_ = cmd.MarkFlagRequired(FlagAddress)
	cmd.Flags().StringP(FlagtokenID, "t", "", "mint token id")
	_ = cmd.MarkFlagRequired(FlagtokenID)

	return cmd
}

func approveERC721Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve",
		Short: "Approve token",
		RunE: func(cmd *cobra.Command, args []string) error {
			erc721CMD, err := setupERC721CMD(ctx)
			if err != nil {
				return err
			}

			spender, err := cmd.Flags().GetString(FlagAddress)
			if err != nil {
				return err
			}
			tokenID, err := cmd.Flags().GetString(FlagtokenID)
			if err != nil {
				return err
			}
			i, ok := new(big.Int).SetString(tokenID, 10)
			if !ok {
				return errors.New("failed to convert")
			}

			tx, err := erc721CMD.Approve(common.HexToAddress(spender), i)
			if err != nil {
				return err
			}

			printTx(tx)
			return nil
		},
	}

	cmd.Flags().StringP(FlagAddress, "a", "", "spender address")
	_ = cmd.MarkFlagRequired(FlagAddress)
	cmd.Flags().StringP(FlagtokenID, "t", "", "token id")
	_ = cmd.MarkFlagRequired(FlagtokenID)

	return cmd
}

func getApprovedERC721Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getApproved",
		Short: "Get owner for token",
		RunE: func(cmd *cobra.Command, args []string) error {
			erc721CMD, err := setupERC721CMD(ctx)
			if err != nil {
				return err
			}

			tokenID, err := cmd.Flags().GetString(FlagtokenID)
			if err != nil {
				return err
			}
			i, ok := new(big.Int).SetString(tokenID, 10)
			if !ok {
				return errors.New("failed to convert")
			}

			result, err := erc721CMD.GetApproved(i)
			if err != nil {
				return err
			}

			fmt.Println(result.Hex())
			return nil
		},
	}

	cmd.Flags().StringP(FlagtokenID, "t", "", "token id")
	_ = cmd.MarkFlagRequired(FlagtokenID)

	return cmd
}

func balanceOfERC721Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balanceOf",
		Short: "Get balance",
		RunE: func(cmd *cobra.Command, args []string) error {
			erc721CMD, err := setupERC721CMD(ctx)
			if err != nil {
				return err
			}

			account, err := cmd.Flags().GetString(FlagAddress)
			if err != nil {
				return err
			}

			result, err := erc721CMD.BalanceOf(common.HexToAddress(account))
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringP(FlagAddress, "a", "", "address")
	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}

func ownerOfERC721Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ownerOf",
		Short: "Get owner",
		RunE: func(cmd *cobra.Command, args []string) error {
			erc721CMD, err := setupERC721CMD(ctx)
			if err != nil {
				return errors.WithMessage(err, "Failed to setup")
			}

			tokenID, err := cmd.Flags().GetString(FlagtokenID)
			if err != nil {
				return err
			}
			i, ok := new(big.Int).SetString(tokenID, 10)
			if !ok {
				return errors.New("failed to convert")
			}

			result, err := erc721CMD.OwnerOf(i)
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringP(FlagtokenID, "t", "", "token id")
	_ = cmd.MarkFlagRequired(FlagtokenID)

	return cmd
}

func printTx(tx *types.Transaction) {
	fmt.Println(tx.Hash().Hex())
}
