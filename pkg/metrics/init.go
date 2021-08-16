package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"mysql_mate_go/pkg/util"
	"time"
)

const (
	namespace = "mysql"
)

func Init() {
	prometheus.Unregister(prometheus.NewGoCollector())
	prometheus.MustRegister()
	util.Schedule(30, time.Second, globalStatusMetrics)
	util.Schedule(30, time.Second, slaveStatusMetrics)
	util.Schedule(30, time.Second, innodbStatusMetrics)
	util.Schedule(30, time.Second, upMetrics)
}
