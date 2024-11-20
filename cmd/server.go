package cmd

import (
	"certgen/internal/server"

	"github.com/spf13/cobra"
)

var (
	ServerCmd = &cobra.Command{
		Use:     "server",
		Short:   "Starts the Certgen server",
		Long:    "Starts the Certgen server",
		Aliases: []string{"serve"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Run()
		},
	}
)
