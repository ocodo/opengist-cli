package cmd

import (
	"fmt"
	"os"

	"github.com/ocodo/opengist-cli/internal/client"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [options] <gist>",
	Short: "Edit an existing gist",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		isPrivate, _ := cmd.Flags().GetBool("private")
		isPublic, _ := cmd.Flags().GetBool("public")

		if isPrivate && isPublic {
			fmt.Fprintln(os.Stderr, "opengist-cli: --private and --public are mutually exclusive")
			os.Exit(2)
		}

		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")

		payload := &client.EditPayload{}

		if cmd.Flags().Changed("title") {
			payload.Title = title
		}
		if cmd.Flags().Changed("description") {
			payload.Description = description
		}

		if isPrivate {
			payload.Visibility = "private"
		} else if isPublic {
			payload.Visibility = "public"
		}

		if payload.Title == "" && payload.Description == "" && payload.Visibility == "" {
			fmt.Fprintln(os.Stderr, "opengist-cli: nothing to edit")
			os.Exit(2)
		}

		apiClient, err := client.NewClient()
		if err != nil {
			return err
		}

		resp, err := apiClient.Edit(args[0], payload)
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
	editCmd.Flags().StringP("title", "t", "", "Set gist title")
	editCmd.Flags().StringP("description", "d", "", "Set gist description")
	editCmd.Flags().Bool("private", false, "Make gist private")
	editCmd.Flags().Bool("public", false, "Make gist public")
	RootCmd.AddCommand(editCmd)
}
