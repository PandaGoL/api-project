package metrics

import (
	"time"
)

// IMetrics - Главный интерфейс метрик.
type IMetrics interface {
	IAppMetrics
	IHTTPMetrics
	IDatabaseMetrics
	IPanic
}

// IAppMetrics - Интерфейс метрик приложения
type IAppMetrics interface {
	SetApplicationInfo(commit, branch, version, buildDate string)
}

// IHTTPMetrics - интерфейс HTTP метрик.
type IHTTPMetrics interface {
	AddHTTPIncomingRequests(route string)
	AddHTTPIncomingResponses(route string, status int, dur time.Duration)
}

// IDatabaseMetrics - интерфейс метрик хранилища.
type IDatabaseMetrics interface {
	AddDBRequests(store, query, status string, dur time.Duration)
}

// IPanic - Интерфейс метрик паник сервиса
type IPanic interface {
	AddPanic(status string)
}
