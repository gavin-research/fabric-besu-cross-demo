package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc20/config"
	erc20mgrtypes "github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20/erc20mgr/types"
	"github.com/spf13/cobra"
)

func erc20Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "erc20",
		Short:                      "ERC20 subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		mintERC20Cmd(ctx),
		approveERC20Cmd(ctx),
		allowanceERC20Cmd(ctx),
		balanceOfERC20Cmd(ctx),
		totalSupplyERC20Cmd(ctx),
	)

	return cmd
}

func mintERC20Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint",
		Short: "Mint token",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ctx.Chain.Connect(); err != nil {
				return err
			}

			receiverMSPID, err := cmd.Flags().GetString(FlagReceiverMSPID)
			if err != nil {
				return err
			}
			receiverClientID, err := cmd.Flags().GetString(FlagReceiverClientID)
			if err != nil {
				return err
			}
			receiver, err := erc20mgrtypes.NewID(erc20mgrtypes.MSPID(receiverMSPID), erc20mgrtypes.ClientID(receiverClientID))
			if err != nil {
				return err
			}
			receiverBytes, err := receiver.Marshal()
			if err != nil {
				return err
			}
			amount, err := cmd.Flags().GetString(FlagAmount)
			if err != nil {
				return err
			}
			mspID := ctx.Config.WalletLabel
			clientID := ctx.ClientID
			id, err := erc20mgrtypes.NewID(erc20mgrtypes.MSPID(mspID), erc20mgrtypes.ClientID(clientID))
			if err != nil {
				return err
			}
			idBytes, err := id.Marshal()
			if err != nil {
				return err
			}
			msg := erc20mgrtypes.NewMsgContractCallTx(
				&erc20mgrtypes.ContractCallRequest{
					Method: "mint",
					Args:   []string{hex.EncodeToString(receiverBytes), amount},
				},
				[]authtypes.AccountID{
					idBytes,
				},
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			_, err = ctx.Chain.SendMsgs([]sdk.Msg{msg})
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().String(FlagReceiverMSPID, "", "msp id")
	_ = cmd.MarkFlagRequired(FlagReceiverMSPID)
	cmd.Flags().String(FlagReceiverClientID, "", "client id")
	_ = cmd.MarkFlagRequired(FlagReceiverClientID)
	cmd.Flags().String(FlagAmount, "", "amount")
	_ = cmd.MarkFlagRequired(FlagAmount)

	return cmd
}

func approveERC20Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve",
		Short: "Approve token",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ctx.Chain.Connect(); err != nil {
				return err
			}

			spenderMSPID, err := cmd.Flags().GetString(FlagSpenderMSPID)
			if err != nil {
				return err
			}
			spenderClientID, err := cmd.Flags().GetString(FlagSpenderClientID)
			if err != nil {
				return err
			}
			spender, err := erc20mgrtypes.NewID(erc20mgrtypes.MSPID(spenderMSPID), erc20mgrtypes.ClientID(spenderClientID))
			if err != nil {
				return err
			}
			spenderBytes, err := spender.Marshal()
			if err != nil {
				return err
			}
			amount, err := cmd.Flags().GetString(FlagAmount)
			if err != nil {
				return err
			}
			mspID := ctx.Config.WalletLabel
			clientID := ctx.ClientID
			id, err := erc20mgrtypes.NewID(erc20mgrtypes.MSPID(mspID), erc20mgrtypes.ClientID(clientID))
			if err != nil {
				return err
			}
			idBytes, err := id.Marshal()
			if err != nil {
				return err
			}

			msg := erc20mgrtypes.NewMsgContractCallTx(
				&erc20mgrtypes.ContractCallRequest{
					Method: "approve",
					Args:   []string{hex.EncodeToString(spenderBytes), amount},
				},
				[]authtypes.AccountID{idBytes},
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			_, err = ctx.Chain.SendMsgs([]sdk.Msg{msg})
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().String(FlagSpenderMSPID, "", "msp id")
	_ = cmd.MarkFlagRequired(FlagSpenderMSPID)
	cmd.Flags().String(FlagSpenderClientID, "", "client id")
	_ = cmd.MarkFlagRequired(FlagSpenderClientID)
	cmd.Flags().String(FlagAmount, "", "amount")
	_ = cmd.MarkFlagRequired(FlagAmount)

	return cmd
}

func allowanceERC20Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allowance",
		Short: "Get allowance",
		RunE: func(cmd *cobra.Command, args []string) error {
			ownerMSPID, err := cmd.Flags().GetString(FlagOwnerMSPID)
			if err != nil {
				return err
			}
			ownerClientID, err := cmd.Flags().GetString(FlagOwnerClientID)
			if err != nil {
				return err
			}
			owner, err := erc20mgrtypes.NewID(erc20mgrtypes.MSPID(ownerMSPID), erc20mgrtypes.ClientID(ownerClientID))
			if err != nil {
				return err
			}
			spenderMSPID, err := cmd.Flags().GetString(FlagSpenderMSPID)
			if err != nil {
				return err
			}
			spenderClientID, err := cmd.Flags().GetString(FlagSpenderClientID)
			if err != nil {
				return err
			}
			spender, err := erc20mgrtypes.NewID(erc20mgrtypes.MSPID(spenderMSPID), erc20mgrtypes.ClientID(spenderClientID))
			if err != nil {
				return err
			}

			res, err := ctx.Chain.QueryAllowance(owner, spender)
			if err != nil {
				return err
			}

			j, err := ctx.Chain.Codec().MarshalJSON(res)
			if err != nil {
				return err
			}
			fmt.Println(string(j))
			return nil
		},
	}

	cmd.Flags().String(FlagOwnerMSPID, "", "msp id")
	_ = cmd.MarkFlagRequired(FlagOwnerMSPID)
	cmd.Flags().String(FlagOwnerClientID, "", "client id")
	_ = cmd.MarkFlagRequired(FlagOwnerClientID)
	cmd.Flags().String(FlagSpenderMSPID, "", "msp id")
	_ = cmd.MarkFlagRequired(FlagSpenderMSPID)
	cmd.Flags().String(FlagSpenderClientID, "", "client id")
	_ = cmd.MarkFlagRequired(FlagSpenderClientID)

	return cmd
}

func balanceOfERC20Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balanceOf",
		Short: "Get balance",
		RunE: func(cmd *cobra.Command, args []string) error {
			ownerMSPID, err := cmd.Flags().GetString(FlagOwnerMSPID)
			if err != nil {
				return err
			}
			ownerClientID, err := cmd.Flags().GetString(FlagOwnerClientID)
			if err != nil {
				return err
			}
			owner, err := erc20mgrtypes.NewID(erc20mgrtypes.MSPID(ownerMSPID), erc20mgrtypes.ClientID(ownerClientID))
			if err != nil {
				return err
			}

			res, err := ctx.Chain.QueryBalanceOf(owner)
			if err != nil {
				return err
			}

			fmt.Println(res.Balance.Amount)
			return nil
		},
	}

	cmd.Flags().String(FlagOwnerMSPID, "", "msp id")
	_ = cmd.MarkFlagRequired(FlagOwnerMSPID)
	cmd.Flags().String(FlagOwnerClientID, "", "client id")
	_ = cmd.MarkFlagRequired(FlagOwnerClientID)

	return cmd
}

func totalSupplyERC20Cmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "totalSupply",
		Short: "Get totalSupply",
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := ctx.Chain.QueryTotalSupply()
			if err != nil {
				return err
			}

			fmt.Println(res.TotalSupply)
			return nil
		},
	}

	return cmd
}
