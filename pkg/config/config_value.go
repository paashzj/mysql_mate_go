package config

import (
	"mysql_mate_go/pkg/util"
)

// mate config
var (
	Cluster    string
	LogFile    string
	ConsoleLog bool
	ListenAddr string
	RemoteMode bool
)

// mysql config
var (
	Password string
	Hostname string
	Username string
)

func init() {
	Cluster = util.GetEnvStr("CLUSTER", "default")
	LogFile = util.GetEnvStr("LOG_FILE", "")
	ConsoleLog = util.GetEnvBool("CONSOLE_LOG", true)
	ListenAddr = util.GetEnvStr("LISTEN_ADDR", "0.0.0.0")
	RemoteMode = util.GetEnvBool("REMOTE_MODE", true)
	Username = util.GetEnvStr("USERNAME", "hzj")
	Password = util.GetEnvStr("PASSWORD", "Mysql@123")
	Hostname = util.GetEnvStr("HOSTNAME", "127.0.0.1:3306")
}
