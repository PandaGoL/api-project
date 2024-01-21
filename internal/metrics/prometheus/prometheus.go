package prometheus

import (
	"strings"

	"github.com/PandaGoL/api-project/internal/metrics/prometheus/common"

	"github.com/prometheus/client_golang/prometheus"
)

// NewMetrics - создает *Base.
func NewMetrics(applicationName string) *Base {
	namespace := strings.ReplaceAll(applicationName, "-", "_")

	m := &Base{
		HTTPMetrics: common.NewHTTPMetrics(namespace),
		DBMetrics:   common.NewDBMetrics(namespace),
	}

	m.applicationInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "application_info",
		Help:      "Информация о приложении",
	}, []string{"commit", "branch", "version", "build_date"})

	prometheus.MustRegister(m.applicationInfo)

	m.panics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:      "panics_counter",
		Namespace: namespace,
		Help:      "Счетчик паник",
	}, []string{"source"},
	)

	prometheus.MustRegister(m.panics)

	return m
}

// Base - структура реализующая интерфейс metrics.IMetrics на prometheus
type Base struct {
	*common.HTTPMetrics
	*common.DBMetrics

	applicationInfo *prometheus.GaugeVec
	panics          *prometheus.CounterVec
}

// SetApplicationInfo - Отправляем информацию о приложении.
func (m *Base) SetApplicationInfo(commit, branch, version, buildDate string) {
	m.applicationInfo.WithLabelValues(commit, branch, version, buildDate).Set(1)
}

// AddPanic - отправляет в метрику паники
func (m *Base) AddPanic(label string) {
	m.panics.WithLabelValues(label).Inc()
}
