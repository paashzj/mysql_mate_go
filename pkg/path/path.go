package path

import (
	"os"
	"path/filepath"
)

// mysql
var (
	MysqlHome = os.Getenv("MYSQL_HOME")
)

// mate
var (
	MysqlMatePath    = filepath.FromSlash(MysqlHome + "/mate")
	MysqlScripts     = filepath.FromSlash(MysqlMatePath + "/scripts")
	MysqlStartScript = filepath.FromSlash(MysqlScripts + "/start-mysql.sh")
)
