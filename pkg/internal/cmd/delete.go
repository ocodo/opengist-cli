package cmd

import (
	"fmt"
	"os"

	"github.com/ocodo/opengist-cli/internal/client"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <gist>",
	Short: "Delete an existing gist",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := client.NewClient()
		if err != nil {
			return err
		}

		gistID := args[0]
		resp, err := apiClient.Delete(gistID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "opengist-cli: request failed: %v\n", err)
			os.Exit(1)
		}

		if resp.IsSuccess() {
			fmt.Printf("Deleted %s\n", gistID)
			return nil
		}

		client.HandleResponse(resp)
		os.Exit(1)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
