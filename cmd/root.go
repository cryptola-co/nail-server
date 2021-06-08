package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "nail-server",
	Short: "nail is an agnostic container orchestration tool",
}

func Execute() error {
	return rootCmd.Execute()
}
