package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	"github.com/datachainlab/cross/x/core/initiator/types"
	xcctypes "github.com/datachainlab/cross/x/core/xcc/types"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc721/config"
	exttypes "github.com/datachainlab/fabric-besu-cross-demo/contracts/extension/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io/fs"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

func crossCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cross",
		Aliases: []string{"cross"},
		Short:   "manage cross tx",
	}

	cmd.AddCommand(
		getAddressCrossCmd(ctx),
		createContractTransactionCrossCmd(ctx),
		callInfoCrossCmd(ctx),
	)

	return cmd
}

func getAddressCrossCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address",
		Short: "Get Cross contract address",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(ctx.Config.CrossModuleAddress)
			return nil
		},
	}

	return cmd
}

func createContractTransactionCrossCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-contract-tx",
		Short: "Create a new contract transaction",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			signer, err := cmd.Flags().GetString(FlagSigner)
			if err != nil {
				return err
			}

			account := authtypes.Account{
				Id:       common.HexToAddress(signer).Bytes(),
				AuthType: authtypes.NewAuthTypeExtension(&exttypes.BesuAuthExtension{}),
			}

			callInfo, err := cmd.Flags().GetString(FlagCallInfo)
			if err != nil {
				return err
			}
			callInfoBytes, err := hex.DecodeString(callInfo)
			if err != nil {
				return err
			}

			initiatorChannel, err := cmd.Flags().GetString(FlagInitiatorChainChannel)
			if err != nil {
				return err
			}
			ci, err := parseChannelInfoFromString(initiatorChannel)
			if err != nil {
				return err
			}
			xcc, err := xcctypes.PackCrossChainChannel(ci)
			if err != nil {
				return err
			}

			cTx := types.ContractTransaction{
				CrossChainChannel: xcc,
				Signers:           []authtypes.Account{account},
				CallInfo:          callInfoBytes,
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

	cmd.Flags().String(FlagInitiatorChainChannel, "", "The channel info: '<channelID>:<portID>'")
	_ = cmd.MarkFlagRequired(FlagInitiatorChainChannel)
	cmd.Flags().String(FlagSigner, "", "signer")
	_ = cmd.MarkFlagRequired(FlagSigner)
	cmd.Flags().String(FlagCallInfo, "", "call info")
	_ = cmd.MarkFlagRequired(FlagCallInfo)
	cmd.Flags().String(flags.FlagOutputDocument, "", "The document will be written to the given file instead of STDOUT")

	return cmd
}

func callInfoCrossCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "call-info",
		Short: "Create a call info for contract transaction",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			contract, err := cmd.Flags().GetString(FlagAddress)
			if err != nil {
				return err
			}
			funcName, err := cmd.Flags().GetString(FlagFunctionName)
			if err != nil {
				return err
			}
			arguments, err := cmd.Flags().GetStringArray(FlagArguments)
			if err != nil {
				return err
			}
			argumentTypes, err := cmd.Flags().GetStringArray(FlagArgumentTypes)
			if err != nil {
				return err
			}
			if len(arguments) != len(argumentTypes) {
				return errors.New("arguments and argumentTypes are different length")
			}

			var convertedArgs []interface{}
			for i, a := range arguments {
				if a != "" {
					t := argumentTypes[i]
					arg, err := toArg(a, t)
					if err != nil {
						return err
					}
					convertedArgs = append(convertedArgs, arg)
				}
			}

			callInfo, err := toCallInfo(common.HexToAddress(contract), funcName, convertedArgs)
			if err != nil {
				return err
			}

			fmt.Println(hex.EncodeToString(callInfo))

			return nil
		},
	}

	cmd.Flags().String(FlagAddress, "", "contract address")
	_ = cmd.MarkFlagRequired(FlagAddress)
	cmd.Flags().String(FlagFunctionName, "", "function name")
	_ = cmd.MarkFlagRequired(FlagFunctionName)
	cmd.Flags().StringArray(FlagArguments, []string{""}, "arguments")
	cmd.Flags().StringArray(FlagArgumentTypes, []string{""}, "argument type")
	return cmd
}

func parseChannelInfoFromString(s string) (*xcctypes.ChannelInfo, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return nil, errors.New("channel format must be follow a format: '<channelID>:<portID>'")
	}
	return &xcctypes.ChannelInfo{Channel: parts[0], Port: parts[1]}, nil
}

func setOutputFile(cmd *cobra.Command) (func(), error) {
	outputDoc, err := cmd.Flags().GetString(flags.FlagOutputDocument)
	if err != nil {
		return func() {}, err
	}
	if outputDoc == "" {
		cmd.SetOut(cmd.OutOrStdout())
		return func() {}, nil
	}

	dir := filepath.Dir(outputDoc)
	if _, err := os.Stat(dir); errors.Is(err, fs.ErrNotExist) {
		if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	fp, err := os.OpenFile(outputDoc, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return func() {}, err
	}

	cmd.SetOut(fp)

	return func() { fp.Close() }, nil
}

type callInfo struct {
	Contract     common.Address
	FunctionName string
	Args         []interface{}
}

func toCallInfo(contract common.Address, funcName string, args []interface{}) ([]byte, error) {
	ci := callInfo{
		Contract:     contract,
		FunctionName: funcName,
		Args:         args,
	}
	return rlp.EncodeToBytes(ci)
}

func toArg(raw string, types string) (interface{}, error) {
	switch types {
	case "string":
		return raw, nil
	case "address":
		return common.HexToAddress(raw), nil
	case "int":
		i, ok := new(big.Int).SetString(raw, 10)
		if !ok {
			return nil, errors.New("failed to convert")
		}
		return i, nil
	default:
		return nil, errors.New("invalid args")
	}
}
