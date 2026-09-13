package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ocodo/opengist-cli/internal/client"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [options] <file ...>",
	Short: "Create a new gist",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, err := client.NewClient()
		if err != nil {
			return err
		}

		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		isPrivate, _ := cmd.Flags().GetBool("private")

		files := make(map[string]client.FileContent)
		for _, filename := range args {
			content, err := os.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "opengist-cli: file not found: %s\n", filename)
				os.Exit(1)
			}
			baseName := filepath.Base(filename)
			files[baseName] = client.FileContent{Content: string(content)}
		}

		payload := &client.AddPayload{
			Files:       files,
			Title:       title,
			Description: description,
		}

		if isPrivate {
			payload.Visibility = "private"
		}

		resp, err := apiClient.Add(payload)
		if err != nil {
			fmt.Fprintf(os.Stderr, "opengist-cli: request failed: %v\n", err)
			os.Exit(1)
		}

		client.HandleResponse(resp)
		if !resp.IsSuccess() {
			os.Exit(1)
		}
		return nil
	},
}

func init() {
	addCmd.Flags().StringP("title", "t", "", "Gist title")
	addCmd.Flags().StringP("description", "d", "", "Gist description")
	addCmd.Flags().Bool("private", false, "Create a private gist")
	RootCmd.AddCommand(addCmd)
}
