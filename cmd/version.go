package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of nail-server",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("nail server application %s -- RELEASE\n", version)
	},
}
