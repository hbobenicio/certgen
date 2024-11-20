package database

import (
	"embed"
	"fmt"
)

//go:embed assets
var assets embed.FS

const (
	CreateTableCertsPath = "assets/001_create_table_certs.sql"
)

func GetCreateTableCertsSql() (string, error) {
	sqlBytes, err := assets.ReadFile(CreateTableCertsPath)
	if err != nil {
		return "", fmt.Errorf("failed to read embed file \"%s\": %w", CreateTableCertsPath, err)
	}

	return string(sqlBytes), nil
}
