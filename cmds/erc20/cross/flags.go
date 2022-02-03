package cross

import "github.com/spf13/cobra"

const (
	flagTxID        = "tx-id"
	flagLightHeight = "light-height"
	flagEthSignKey  = "eth-sign-key"
)

func txIdFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().String(flagTxID, "", "hex encoding of the TxID")
	_ = cmd.MarkFlagRequired(flagTxID)
	return cmd
}

func lightHeightFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().Uint64(flagLightHeight, 0, "Latest light height")
	return cmd
}

func ethSignKeyFlag(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().String(flagEthSignKey, "", "the Ethereum Chain private key used by the importer for signing")
	return cmd
}
