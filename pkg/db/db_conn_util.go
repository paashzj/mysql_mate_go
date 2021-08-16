package db

import (
	"fmt"
	"mysql_mate_go/pkg/config"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Username, config.Password, config.Hostname, dbName)
}
