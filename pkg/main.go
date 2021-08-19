package main

import (
	"flag"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"mysql_mate_go/pkg/config"
	_ "mysql_mate_go/pkg/config"
	"mysql_mate_go/pkg/metrics"
	"mysql_mate_go/pkg/mysql"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
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
		logs.Info("not remote mode, start redis server")
		mysql.Start()
	} else {
		logs.Info("remote mode")
	}
	metrics.Init()
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	http.Handle("/", r)
	err := http.ListenAndServe(config.ListenAddr+":8080", nil)
	if err != nil {
		logs.Error("open http server failed")
		return
	}
	logs.Info("mysql adapter started")
}
