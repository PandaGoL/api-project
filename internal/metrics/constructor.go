package metrics

import (
	"github.com/PandaGoL/api-project/internal/metrics/prometheus"
	log "github.com/sirupsen/logrus"
)

func New(namespace string) Metrics {

	metrics := prometheus.NewMetrics(namespace)

	log.Warn("Metrics initialization successfully")

	return metrics
}
