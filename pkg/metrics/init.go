package metrics

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/prometheus/client_golang/prometheus"
	"mysql_mate_go/pkg/util"
	"time"
)

const (
	namespace = "mysql"
)

func Init() {
	prometheus.MustRegister()
	err := prometheus.Register(newGlobalStatusCounterCollector())
	if err != nil {
		logs.Error("register global status collector failed")
	}
	util.Schedule(30, time.Second, globalStatusMetrics)
	util.Schedule(30, time.Second, slaveStatusMetrics)
	util.Schedule(30, time.Second, innodbStatusMetrics)
	util.Schedule(30, time.Second, upMetrics)
}
