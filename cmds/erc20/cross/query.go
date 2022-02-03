package cross

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	"github.com/datachainlab/cross/x/core/initiator/types"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc20/config"
	erc20mgrtypes "github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20/erc20mgr/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/fs"
	"os"
	"path/filepath"
)

func GetCreateContractTransaction(ctx *config.Context) *cobra.Command {
	const (
		flagSignerMspID    = "signer-msp-id"
		flagSignerClientID = "signer-client-id"
		flagCallInfo       = "call-info"
		flagOutput         = "output"
	)

	cmd := &cobra.Command{
		Use:   "create-contract-tx",
		Short: "Create a new contract transaction",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var anyXCC *codectypes.Any

			// Query self-XCC to chaincode
			res, err := ctx.Chain.QuerySelfXCC()
			if err != nil {
				return err
			}
			anyXCC = res.Xcc

			mspID := viper.GetString(flagSignerMspID)
			clientID := viper.GetString(flagSignerClientID)
			id, err := erc20mgrtypes.NewID(erc20mgrtypes.MSPID(mspID), erc20mgrtypes.ClientID(clientID))
			if err != nil {
				return err
			}
			idBytes, err := id.Marshal()
			if err != nil {
				return err
			}
			signer := authtypes.Account{
				Id:       idBytes,
				AuthType: authtypes.NewAuthTypeExtension(&erc20mgrtypes.FabricAuthExtension{}),
			}

			callInfo := []byte(viper.GetString(flagCallInfo))
			cTx := types.ContractTransaction{
				CrossChainChannel: anyXCC,
				Signers:           []authtypes.Account{signer},
				CallInfo:          callInfo,
			}
			// prepare output document
			closeFunc, err := setOutputFile(cmd)
			if err != nil {
				return err
			}
			defer closeFunc()

			bz, err := ctx.Codec.MarshalJSON(&cTx)
			if err != nil {
				return err
			}

			if _, err := cmd.OutOrStdout().Write(bz); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().String(flagSignerMspID, "", "The mspId to sign to initiateTx")
	_ = cmd.MarkFlagRequired(flagSignerMspID)

	cmd.Flags().String(flagSignerClientID, "", "The clientId to sign to initiateTx")
	_ = cmd.MarkFlagRequired(flagSignerClientID)

	cmd.Flags().String(flagCallInfo, "", "A contract call info")
	_ = cmd.MarkFlagRequired(flagCallInfo)

	cmd.Flags().String(flags.FlagOutputDocument, "", "The document will be written to the given file instead of STDOUT")
	cmd.Flags().StringP(flagOutput, "o", "text", "Output format (text|json)")

	ethSignKeyFlag(cmd)

	return cmd
}

func setOutputFile(cmd *cobra.Command) (func(), error) {
	outputDoc, err := cmd.Flags().GetString(flags.FlagOutputDocument)
	if err != nil {
		return nil, err
	}
	if outputDoc == "" {
		cmd.SetOut(cmd.OutOrStdout())
		return nil, nil
	}

	dir := filepath.Dir(outputDoc)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
			return nil, err
		}
	}

	fp, err := os.OpenFile(outputDoc, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}

	cmd.SetOut(fp)

	return func() { fp.Close() }, nil
}

func QueryTxAuthStateCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx-auth-state [tx-id]",
		Short: "Query the state of a client in a given path",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if txId, err := hex.DecodeString(args[0]); err != nil {
				return err
			} else if res, err := ctx.Chain.QueryTxAuthState(txId); err != nil {
				return err
			} else if bz, err := ctx.Chain.Codec().MarshalJSON(res); err != nil {
				return err
			} else {
				fmt.Println(string(bz))
			}

			return nil
		},
	}

	return cmd
}
