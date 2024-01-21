package common

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// NewHTTPMetrics - создает *HTTPMetrics.
func NewHTTPMetrics(namespace string) *HTTPMetrics {
	m := &HTTPMetrics{}

	m.httpIncomingRequestsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "http_incoming_requests_counter",
		Help:      "Счетчик входящих запросов HTTP",
	}, []string{"route"})

	m.httpIncomingResponsesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "http_incoming_responses_counter",
		Help:      "Счетчик ответов на запросы HTTP",
	}, []string{"status", "route"})

	m.httpIncomingResponsesDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "http_incoming_responses_duration",
		Help:      "Длительность ответов на запросы HTTP",
	}, []string{"status", "route"})

	prometheus.MustRegister(
		m.httpIncomingRequestsCounter,
		m.httpIncomingResponsesCounter,
		m.httpIncomingResponsesDuration)

	return m
}

// HTTPMetrics - структура метрик http запросов.
type HTTPMetrics struct {
	httpIncomingRequestsCounter   *prometheus.CounterVec
	httpIncomingResponsesCounter  *prometheus.CounterVec
	httpIncomingResponsesDuration *prometheus.HistogramVec
}

// AddHTTPIncomingRequests - Считаем количество входящих запросов к Http API.
func (m *HTTPMetrics) AddHTTPIncomingRequests(route string) {
	m.httpIncomingRequestsCounter.WithLabelValues(route).Inc()
}

// AddHTTPIncomingResponses - Считаем количество ответов на входящие запросы к Http API и длительность их исполнения.
func (m *HTTPMetrics) AddHTTPIncomingResponses(route string, status int, dur time.Duration) {
	strStatus := strconv.Itoa(status)
	m.httpIncomingResponsesCounter.WithLabelValues(route, strStatus).Inc()
	m.httpIncomingResponsesDuration.WithLabelValues(route, strStatus).Observe(dur.Seconds() * 1000)
}
