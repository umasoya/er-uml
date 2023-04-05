package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version %s\n", "0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
