package cmd

import (
	"encoding/json"
	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"io/ioutil"
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/codec"
	cross "github.com/datachainlab/cross/x/core"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc20/config"
	erc20types "github.com/datachainlab/fabric-besu-cross-demo/contracts/erc20/erc20mgr/types"
	"github.com/hyperledger-labs/yui-relayer/chains/fabric"
	msppb "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	homePath    string
	defaultHome = os.ExpandEnv("$HOME/.erc20cli")
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "erc20cli",
	Short: "This provides commands for cross on a Fabric chain",
}

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.SilenceUsage = true

	// Register top level flags --home and --debug
	rootCmd.PersistentFlags().StringVar(&homePath, flags.FlagHome, defaultHome, "set home directory")
	if err := viper.BindPFlag(flags.FlagHome, rootCmd.PersistentFlags().Lookup(flags.FlagHome)); err != nil {
		panic(err)
	}

	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	fabric.RegisterInterfaces(interfaceRegistry)
	erc20types.RegisterInterfaces(interfaceRegistry)
	cryptotypes.RegisterInterfaces(interfaceRegistry)
	cross.AppModuleBasic{}.RegisterInterfaces(interfaceRegistry)

	ctx := &config.Context{Config: &fabric.ChainConfig{}, Codec: cdc}

	// Register subcommands
	rootCmd.AddCommand(
		configCmd(ctx),
		crossCmd(ctx),
		ibcCmd(ctx),
		erc20Cmd(ctx),
		fabricCmd(ctx),
	)

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		return initConfig(ctx, rootCmd)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
// config directory as below
// home
// ├── configPath
// │   └── configFile
// └── walletPath
//     └── walletFile (<wallet_label>.id)
func initConfig(ctx *config.Context, cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(flags.FlagHome)
	if err != nil {
		return err
	}

	// Set config to context
	cfgPath := path.Join(home, config.ConfigPath, config.ConfigFile)
	if _, err := os.Stat(cfgPath); err == nil {
		viper.SetConfigFile(cfgPath)
		if err := viper.ReadInConfig(); err == nil {
			file, err := ioutil.ReadFile(viper.ConfigFileUsed())
			if err != nil {
				return err
			}

			err = ctx.Codec.UnmarshalJSON(file, ctx.Config)
			if err != nil {
				return err
			}

			iChain, err := ctx.Config.Build()
			if err != nil {
				return err
			}

			ctx.Chain = &config.Chain{
				Chain: iChain.(*fabric.Chain),
			}
			if err := ctx.Chain.Init(home, 0, ctx.Codec, false); err != nil {
				return err
			}

			clientID, err := clientIDFromWallet(home, ctx.Config.WalletLabel)
			if err != nil {
				return err
			}
			ctx.ClientID = clientID
		}
	}
	return nil
}

type chaincodeStub struct {
	Creator []byte
}

func (c *chaincodeStub) GetCreator() ([]byte, error) {
	return c.Creator, nil
}

func clientIDFromWallet(home, walletLabel string) (string, error) {
	walletFile := walletLabel + config.WalletSuffix
	walletPath := path.Join(home, config.WalletPath, walletFile)
	file, err := ioutil.ReadFile(walletPath)
	if err != nil {
		return "", err
	}

	type credentials struct {
		Certificate string `json:"certificate"`
	}
	type jsonResponse struct {
		Credentials credentials `json:"credentials"`
	}
	jsonRes := jsonResponse{}
	err = json.Unmarshal(file, &jsonRes)
	if err != nil {
		return "", err
	}

	sid := &msppb.SerializedIdentity{
		Mspid:   walletLabel,
		IdBytes: []byte(jsonRes.Credentials.Certificate),
	}

	encoded, err := proto.Marshal(sid)
	if err != nil {
		return "", err
	}
	stub := &chaincodeStub{Creator: encoded}
	client, err := cid.New(stub)
	if err != nil {
		return "", err
	}
	return client.GetID()
}
