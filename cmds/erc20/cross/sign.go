package cross

import (
	"encoding/hex"
	erc20mgrtypes "github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20/erc20mgr/types"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	crosstypes "github.com/datachainlab/cross/x/core/auth/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc20/config"
	exttypes "github.com/datachainlab/fabric-besu-cross-demo/contracts/extension/types"
)

func NewSignTxCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ext-sign-tx",
		Short: "Sign the cross-chain transaction on other chain via the chain",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			txID, err := hex.DecodeString(viper.GetString(flagTxID))
			if err != nil {
				return err
			}

			ethSignKey := viper.GetString(flagEthSignKey)
			signer, err := getSigner(ctx, ethSignKey)
			if err != nil {
				return err
			}
			msg := &crosstypes.MsgExtSignTx{
				TxID:    txID,
				Signers: []authtypes.Account{signer},
			}

			var res []byte
			if len(ethSignKey) == 0 {
				if err := ctx.Chain.Connect(); err != nil {
					return err
				} else if res, err = ctx.Chain.SendMsgs([]sdk.Msg{msg}); err != nil {
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

			return nil
		},
	}

	txIdFlag(cmd)
	ethSignKeyFlag(cmd)

	return cmd
}

func getSigner(ctx *config.Context, ethSignKey string) (authtypes.Account, error) {
	var signer authtypes.Account
	if len(ethSignKey) == 0 {
		// when ethSignKey is not specified, ID (with MSPId and ClientID) is included in signers
		id, err := erc20mgrtypes.NewID(erc20mgrtypes.MSPID(ctx.Config.WalletLabel), erc20mgrtypes.ClientID(ctx.ClientID))
		if err != nil {
			return authtypes.Account{}, err
		}
		idBytes, err := id.Marshal()
		if err != nil {
			return authtypes.Account{}, err
		}
		signer = authtypes.Account{
			Id:       idBytes,
			AuthType: authtypes.NewAuthTypeExtension(&erc20mgrtypes.FabricAuthExtension{}),
		}
	} else {
		addr, err := hexToEthereumAddress(ethSignKey)
		if err != nil {
			return authtypes.Account{}, err
		}
		signer = authtypes.Account{
			Id:       addr,
			AuthType: authtypes.NewAuthTypeExtension(&exttypes.BesuAuthExtension{}),
		}
	}
	return signer, nil
}
