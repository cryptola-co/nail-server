package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start nail server",
	Run:   startNailServer,
}

func startNailServer(cmd *cobra.Command, args []string) {

}
