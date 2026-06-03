package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dead-link-doctor [path]",
	Short: "Fast dead-link checker for CI pipelines",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) > 0 { path = args[0] }
		fmt.Printf("Checking links in %s...\n", path)
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
