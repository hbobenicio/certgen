// Package certgen implements an REST API capable of generate self-signed certificates for development purposes.
package main

import (
	"certgen/cmd"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
