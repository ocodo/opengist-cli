package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/ocodo/opengist-cli/internal/client"
	"github.com/spf13/cobra"
)

func strWidth(s string) int {
	w := runewidth.StringWidth(s)
	if w >= 0 {
		return w
	}
	return len(s)
}

func pad(s string, width int) string {
	return s + strings.Repeat(" ", width-strWidth(s))
}

var listCmd = &cobra.Command{
	Use:   "list [options] [username]",
	Short: "List gists",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		page, _ := cmd.Flags().GetInt("page")
		perPage, _ := cmd.Flags().GetInt("per-page")
		markdown, _ := cmd.Flags().GetBool("markdown")

		if page < 1 {
			fmt.Fprintln(os.Stderr, "opengist-cli: --page must be at least 1")
			os.Exit(2)
		}
		if perPage < 1 {
			fmt.Fprintln(os.Stderr, "opengist-cli: --per-page must be at least 1")
			os.Exit(2)
		}

		var username string
		if len(args) > 0 {
			username = args[0]
		}

		apiClient, err := client.NewClient()
		if err != nil {
			return err
		}

		gists, resp, err := apiClient.List(username, page, perPage)
		if err != nil {
			fmt.Fprintf(os.Stderr, "opengist-cli: request failed: %v\n", err)
			os.Exit(1)
		}

		if !resp.IsSuccess() {
			client.HandleResponse(resp)
			os.Exit(1)
		}

		if len(gists) == 0 {
			fmt.Println("No gists found.")
			return nil
		}

		if markdown {
			renderMarkdownTable(gists)
		} else {
			renderTextList(gists)
		}
		return nil
	},
}

func renderTextList(gists []client.Gist) {
	titleW := strWidth("title")
	urlW := strWidth("url")

	for _, g := range gists {
		u := g.HTMLURL
		if u == "" {
			u = g.URL
		}
		if w := strWidth(g.Title); w > titleW {
			titleW = w
		}
		if w := strWidth(u); w > urlW {
			urlW = w
		}
	}

	fmt.Printf("%s %s\n\n", pad("title", titleW), pad("url", urlW))
	for _, g := range gists {
		u := g.HTMLURL
		if u == "" {
			u = g.URL
		}
		fmt.Printf("%s %s\n", pad(g.Title, titleW), pad(u, urlW))
	}
}

func renderMarkdownTable(gists []client.Gist) {
	titleW := strWidth("title")
	urlW := strWidth("url")

	for _, g := range gists {
		u := g.HTMLURL
		if u == "" {
			u = g.URL
		}
		if w := strWidth(g.Title); w > titleW {
			titleW = w
		}
		if w := strWidth(u); w > urlW {
			urlW = w
		}
	}

	fmt.Printf("| %s | %s |\n", pad("title", titleW), pad("url", urlW))
	fmt.Println("|-|-|")
	for _, g := range gists {
		u := g.HTMLURL
		if u == "" {
			u = g.URL
		}
		fmt.Printf("| %s | %s |\n", pad(g.Title, titleW), pad(u, urlW))
	}
}

func init() {
	listCmd.Flags().Int("page", 1, "Page number")
	listCmd.Flags().Int("per-page", 30, "Number of gists per page")
	listCmd.Flags().Bool("markdown", false, "Format list as markdown table")
	RootCmd.AddCommand(listCmd)
}
