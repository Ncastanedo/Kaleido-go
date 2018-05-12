package cmd

import (
	"fmt"
	"os"

	"github.com/kaleido-io/kaleido-go/pkg/kldexerciser"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func initLogging(debugLevel int) {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	switch debugLevel {
	case 0:
		log.SetLevel(log.ErrorLevel)
		break
	case 1:
		log.SetLevel(log.InfoLevel)
		break
	default:
		log.SetLevel(log.DebugLevel)
		break
	}
	log.Debug("Debug level ", debugLevel)
}

var exerciser kldexerciser.Exerciser

func init() {
	cmd.Flags().IntVarP(&exerciser.DebugLevel, "debug", "d", 1, "0=error, 1=info, 2=debug")
	cmd.Flags().StringVarP(&exerciser.URL, "url", "u", "", "JSON/RPC URL for the Ethereum node: https://user:pass@xyz-rpc.kaleido.io")
	cmd.Flags().StringVarP(&exerciser.SolidityFile, "file", "f", "", "Solidity smart contract to compile (and deploy if no contract address supplied)")
	cmd.Flags().StringVarP(&exerciser.Contract, "contract", "c", "", "Pre-deployed contract address (with or without 0x prefix)")
	cmd.Flags().StringVarP(&exerciser.Method, "method", "m", "", "Method in the contract to invoke")
	cmd.Flags().StringArrayVarP(&exerciser.Args, "args", "x", []string{}, "String arguments to pass to contract method")
	cmd.Flags().IntVarP(&exerciser.Txns, "transactions", "t", 1, "Count of transactions to run per worker (0=infinite)")
	cmd.Flags().IntVarP(&exerciser.Workers, "workers", "w", 1, "Number of workers to run")
	cmd.Flags().BoolVarP(&exerciser.ExternalSign, "extsign", "e", true, "Sign externally with generated accounts")
	cmd.Flags().Uint64VarP(&exerciser.ChainID, "chainid", "i", 0, "Chain ID - required for external signing")
	cmd.Flags().StringArrayVarP(&exerciser.Accounts, "accounts", "a", []string{}, "Account addresses - 1/worker needed for internal geth client signing")
	cmd.MarkFlagRequired("url")
}

var cmd = &cobra.Command{
	Use:   "kaleido-go",
	Short: "Sample exerciser for Ethereum permissioned chains - from Kaleido",
	Run: func(cmd *cobra.Command, args []string) {
		initLogging(exerciser.DebugLevel)
		if err := exerciser.Start(); err != nil {
			log.Error("Exerciser Start: ", err)
			os.Exit(1)
		}
	},
}

// Execute is called by the main method of the package
func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
