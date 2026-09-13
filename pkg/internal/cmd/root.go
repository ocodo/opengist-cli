package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "opengist-cli",
	Short: "CLI for Opengist",
	Long: `CLI for Opengist instance. Requires OPENGIST_CLI_TOKEN and OPENGIST_CLI_URL environment variables to be set.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "opengist-cli: %v\n", err)
		os.Exit(1)
	}
}
