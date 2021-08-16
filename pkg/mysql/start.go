package mysql

import (
	"github.com/beego/beego/v2/core/logs"
	"mysql_mate_go/pkg/path"
	"os/exec"
)

func Start() {
	startMysql()
}

func startMysql() {
	command := exec.Command("/bin/bash", path.MysqlStartScript)
	err := command.Start()
	if err != nil {
		logs.Error("start mysql server failed")
	}
	command.Wait()
}