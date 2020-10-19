package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "srectl",
		Short: "A tool to calculate slo parameter",
		Long:  "Work in progress",
	}
)

// Execute executes the root command srectl e adds its subcommands
func Execute() {

	rootCmd.AddCommand(burnRateCmd)
	rootCmd.Execute()
}
