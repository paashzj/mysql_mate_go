package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"mysql_mate_go/pkg/config"
	"mysql_mate_go/pkg/db"
)

func upMetrics() {
	result := db.SelectOne()
	if result {
		upGauge.Set(1)
	} else {
		upGauge.Set(0)
	}
}

var (
	upGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "mysql_up",
		Help: "is mysql running",
		ConstLabels: map[string]string{
			"cluster": config.Cluster,
		},
	})
)
