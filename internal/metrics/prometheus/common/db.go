package common

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// NewDBMetrics - создает *dbMetrics.
func NewDBMetrics(namespace string) *DBMetrics {
	m := &DBMetrics{}

	m.storeRequestsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "store_requests_counter",
		Help:      "Счетчик запросов к хранилищу",
	}, []string{"type", "query", "status"})

	m.storeRequestsDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "store_requests_duration",
		Help:      "Длительность запросов к хранилищу",
	}, []string{"type", "query", "status"})

	prometheus.MustRegister(
		m.storeRequestsCounter,
		m.storeRequestsDuration)

	return m
}

// DBMetrics - структура метрик хранилища.
type DBMetrics struct {
	storeRequestsCounter  *prometheus.CounterVec
	storeRequestsDuration *prometheus.HistogramVec
}

// AddDBRequests - Считаем количество запросов к базам данных и длительность их исполнения.
func (m *DBMetrics) AddDBRequests(store, query, status string, dur time.Duration) {
	m.storeRequestsCounter.WithLabelValues(store, query, status).Inc()
	m.storeRequestsDuration.WithLabelValues(store, query, status).Observe(dur.Seconds() * 1000)
}
