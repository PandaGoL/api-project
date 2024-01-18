package metrics

import (
	"time"
)

// IMetrics - Главный интерфейс метрик.
type Metrics interface {
	AppMetrics
	HTTPMetrics
	DatabaseMetrics
	Panic
}

// IAppMetrics - Интерфейс метрик приложения
type AppMetrics interface {
	SetApplicationInfo(commit, branch, version, buildDate string)
}

// IHTTPMetrics - интерфейс HTTP метрик.
type HTTPMetrics interface {
	AddHTTPIncomingRequests(route string)
	AddHTTPIncomingResponses(route string, status int, dur time.Duration)
}

// IDatabaseMetrics - интерфейс метрик хранилища.
type DatabaseMetrics interface {
	AddDBRequests(store, query, status string, dur time.Duration)
}

// IPanic - Интерфейс метрик паник сервиса
type Panic interface {
	AddPanic(status string)
}
