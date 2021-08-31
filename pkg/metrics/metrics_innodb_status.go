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
	performance = "performance"
)

// innodbStatusMetrics 参考官方的mysql-exporter指标命名
// example: mysql_global_status_aborted_clients
func innodbStatusMetrics() {
	status := db.GetPerformance()
	innodbBufferPoolPagesFree.Set(util.ConvStr2Float(status["Innodb_buffer_pool_pages_free"]))
	innodbBufferPoolPagesTotal.Set(util.ConvStr2Float(status["Innodb_buffer_pool_pages_total"]))
}

var (
	innodbBufferPoolPagesFree = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, performance, "innodb_buffer_pool_pages_free"),
		Help: "seconds_behind_master",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	innodbBufferPoolPagesTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, performance, "slave_sql_running_value"),
		Help: "slave_sql_running_value",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
)
