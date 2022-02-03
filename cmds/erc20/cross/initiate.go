package cross

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	"github.com/datachainlab/cross/x/core/initiator/types"
	txtypes "github.com/datachainlab/cross/x/core/tx/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc20/config"
)

// NewCreateInitiateTxCmd returns the command to create a NewMsgInitiateTx transaction
func NewCreateInitiateTxCmd(ctx *config.Context) *cobra.Command {
	const (
		flagContractTransactions = "contract-txs"
	)

	cmd := &cobra.Command{
		Use:   "create-initiate-tx",
		Short: "Create a NewMsgInitiateTx transaction for a simple commit",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctxs, err := readContractTransactions(ctx.Chain.Codec(), viper.GetStringSlice(flagContractTransactions))
			if err != nil {
				return err
			}

			msg := types.NewMsgInitiateTx(
				nil,
				ctx.Config.ChainId,
				uint64(time.Now().Unix()),
				txtypes.COMMIT_PROTOCOL_SIMPLE,
				ctxs,
				// TODO: Fix to ZeroHeight but currently scenario is failure.
				clienttypes.NewHeight(0, 10000),
				0,
			)

			// prepare output document
			closeFunc, err := setOutputFile(cmd)
			if err != nil {
				return err
			}
			defer closeFunc()

			bz, err := ctx.Codec.MarshalJSON(msg)
			if err != nil {
				return err
			}

			if _, err := cmd.OutOrStdout().Write(bz); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringSlice(flagContractTransactions, nil, "A file path to includes a contract transaction")
	_ = cmd.MarkFlagRequired(flagContractTransactions)

	cmd.Flags().String(flags.FlagOutputDocument, "", "The document will be written to the given file instead of STDOUT")

	return cmd
}

func readContractTransactions(m codec.JSONCodec, pathList []string) ([]types.ContractTransaction, error) {
	var cTxs []types.ContractTransaction
	for _, path := range pathList {
		bz, err := ioutil.ReadFile(filepath.Clean(path))
		if err != nil {
			return nil, err
		}
		var cTx types.ContractTransaction
		if err := m.UnmarshalJSON(bz, &cTx); err != nil {
			return nil, err
		}
		cTxs = append(cTxs, cTx)
	}
	return cTxs, nil
}

// NewSendInitiateTxCmd returns the command to send a NewMsgInitiateTx transaction
func NewSendInitiateTxCmd(ctx *config.Context) *cobra.Command {
	const (
		flagInitiateTx = "initiate-tx"
	)

	cmd := &cobra.Command{
		Use:   "send-initiate-tx",
		Short: "Send a NewMsgInitiateTx transaction for a simple commit",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ctx.Chain.Connect(); err != nil {
				return err
			}

			msg, err := readInitiateTx(ctx.Chain.Codec(), viper.GetString(flagInitiateTx))
			if err != nil {
				return err
			}

			ethSignKey := viper.GetString(flagEthSignKey)
			signer, err := getSigner(ctx, ethSignKey)
			if err != nil {
				return err
			}

			msg.Signers = []authtypes.Account{signer}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			var res []byte
			if len(ethSignKey) == 0 {
				if res, err = ctx.Chain.SendMsgs([]sdk.Msg{msg}); err != nil {
					return err
				}
			} else {
				if privKey, err := hexToSecp256k1PrivKey(ethSignKey); err != nil {
					return err
				} else if txBytes, err := buildTx(ctx.Chain.Codec().InterfaceRegistry(), privKey, msg); err != nil {
					return err
				} else {
					res, err = ctx.Chain.Contract().SubmitTransaction("handleTx", string(txBytes))
					if err != nil {
						return err
					}
					log.Printf("fabric.SendMsgs.result: res='%v' err='%v'", res, err)
				}
			}

			if err := ctx.Chain.OutputTxIDFromEvent(res); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().String(flagInitiateTx, "", "A file path to includes a initiate-tx")
	_ = cmd.MarkFlagRequired(flagInitiateTx)

	cmd.Flags().String(flags.FlagOutputDocument, "", "The document will be written to the given file instead of STDOUT")

	ethSignKeyFlag(cmd)

	return cmd
}

func readInitiateTx(m codec.JSONCodec, path string) (*types.MsgInitiateTx, error) {
	bz, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	var iTx types.MsgInitiateTx
	if err := m.UnmarshalJSON(bz, &iTx); err != nil {
		return nil, err
	}

	return &iTx, nil
}
