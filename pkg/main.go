package main

import (
	"flag"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mysql_mate_go/pkg/api"
	"mysql_mate_go/pkg/config"
	_ "mysql_mate_go/pkg/config"
	"mysql_mate_go/pkg/metrics"
	"mysql_mate_go/pkg/mysql"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logs.EnableFuncCallDepth(true)
	flag.Parse()
	if config.ConsoleLog {
		logs.SetLogger(logs.AdapterConsole)
	}
	if config.LogFile != "" {
		strings := `{"filename":"` + config.LogFile + `", "level":6}`
		logs.SetLogger(logs.AdapterFile, strings)
	}
	if !config.RemoteMode {
		logs.Info("not remote mode, start mysql server")
		mysql.Start()
	} else {
		logs.Info("remote mode")
	}
	metrics.Init()
	router := gin.Default()
	router.GET("/metrics", prometheusHandler())
	router.GET("/v1/mysql/global-status", api.GlobalStatusHandler)
	router.GET("/v1/mysql/slave-status", api.SlaveStatusHandler)
	router.GET("/v1/mysql/memory/global", api.SlaveStatusHandler)
	logs.Info("mysql adapter started")
	err := router.Run(":31009")
	if err != nil {
		logs.Error("open http server failed")
		return
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
