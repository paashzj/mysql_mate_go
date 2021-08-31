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
	globalStatus = "global_status"
)

// globalStatusMetrics 参考官方的mysql-exporter指标命名
// example: mysql_global_status_aborted_clients
func globalStatusMetrics() {
	status := db.GetGlobalStatus()
	statusAbortedClients.Set(util.ConvStr2Float(status["Aborted_clients"]))
	statusThreadsCreated.Set(util.ConvStr2Float(status["Threads_created"]))
	statusThreadsConnected.Set(util.ConvStr2Float(status["Threads_connected"]))
	statusMaxUsedConnections.Set(util.ConvStr2Float(status["Max_used_connections"]))
	statusConnections.Set(util.ConvStr2Float(status["Connections"]))
	statusDelayedWrites.Set(util.ConvStr2Float(status["Delayed_writes"]))
}

var (
	statusAbortedClients = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, globalStatus, "aborted_clients"),
		Help: "status_aborted_clients",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	statusThreadsCreated = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, globalStatus, "threads_created"),
		Help: "threads_created",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	statusThreadsConnected = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, globalStatus, "threads_connected"),
		Help: "threads_connected",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	statusMaxUsedConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, globalStatus, "max_used_connections"),
		Help: "max_used_connections",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	statusConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, globalStatus, "connections"),
		Help: "connections",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
	statusDelayedWrites = promauto.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, globalStatus, "delayed_writes"),
		Help: "delayed_writes",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
)

type globalStatusCounterCollector struct {
	insertCounterDesc *prometheus.Desc
	deleteCounterDesc *prometheus.Desc
	updateCounterDesc *prometheus.Desc
	selectCounterDesc *prometheus.Desc
}

func (c *globalStatusCounterCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.insertCounterDesc
	ch <- c.deleteCounterDesc
	ch <- c.updateCounterDesc
	ch <- c.selectCounterDesc
}

func (c *globalStatusCounterCollector) Collect(ch chan<- prometheus.Metric) {
	status := db.GetGlobalStatus()
	ch <- prometheus.MustNewConstMetric(c.insertCounterDesc, prometheus.CounterValue, util.ConvStr2Float(status["Com_insert"]))
	ch <- prometheus.MustNewConstMetric(c.deleteCounterDesc, prometheus.CounterValue, util.ConvStr2Float(status["Com_delete"]))
	ch <- prometheus.MustNewConstMetric(c.updateCounterDesc, prometheus.CounterValue, util.ConvStr2Float(status["Com_update"]))
	ch <- prometheus.MustNewConstMetric(c.selectCounterDesc, prometheus.CounterValue, util.ConvStr2Float(status["Com_select"]))
}

func newGlobalStatusCounterCollector() *globalStatusCounterCollector {
	return &globalStatusCounterCollector{
		insertCounterDesc: prometheus.NewDesc(prometheus.BuildFQName(namespace, globalStatus, "commands_insert_total"),
			"commands_insert_total", nil, map[string]string{
				"cluster": config.Cluster,
			}),
		deleteCounterDesc: prometheus.NewDesc(prometheus.BuildFQName(namespace, globalStatus, "commands_delete_total"),
			"commands_delete_total", nil, map[string]string{
				"cluster": config.Cluster,
			}),
		updateCounterDesc: prometheus.NewDesc(prometheus.BuildFQName(namespace, globalStatus, "commands_update_total"),
			"commands_update_total", nil, map[string]string{
				"cluster": config.Cluster,
			}),
		selectCounterDesc: prometheus.NewDesc(prometheus.BuildFQName(namespace, globalStatus, "commands_select_total"),
			"commands_select_total", nil, map[string]string{
				"cluster": config.Cluster,
			}),
	}
}
