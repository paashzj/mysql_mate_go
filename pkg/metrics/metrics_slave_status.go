package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"mysql_mate_go/pkg/config"
	"mysql_mate_go/pkg/db"
	"mysql_mate_go/pkg/util"
)

const (
	// Subsystem.
	slaveStatus = "slave_status"
)

// slaveStatusMetrics 参考官方的mysql-exporter指标命名
// example: mysql_global_status_aborted_clients
func slaveStatusMetrics() {
	status := db.GetSlaveStatus()
	slaveStatusSecondsBehindMaster.Set(util.ConvStr2Float(status.SecondsBehindMaster))
	slaveStatusSlaveSqlRunningValue.Set(promConvMysqlStr2Bool(status.SlaveSqlRunning))
	slaveStatusSlaveIoRunningValue.Set(promConvMysqlStr2Bool(status.SlaveIoRunning))
}

var (
	slaveStatusSecondsBehindMaster = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, slaveStatus, "seconds_behind_master"),
		Help: "seconds_behind_master",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	slaveStatusSlaveSqlRunningValue = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, slaveStatus, "slave_sql_running_value"),
		Help: "slave_sql_running_value",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	slaveStatusSlaveIoRunningValue = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, slaveStatus, "slave_io_running_value"),
		Help: "slave_io_running_value",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
)
