package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "certgen",
		Short: "Generate, Monitor and Analyze certificates",
		Long:  "Generate, Monitor and Analyze certificates",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	RootCmd.AddCommand(ServerCmd)
}
