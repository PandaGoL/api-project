package recovery

import (
	"runtime/debug"

	"github.com/PandaGoL/api-project/internal/metrics"
	"github.com/sirupsen/logrus"
)

var (
	panicSignal = make(chan bool, 10)
)

// module - модуль восстановления
type module struct {
	signal     chan bool
	picnickers []Panickier
	metrics    metrics.Metrics
}

// CreateRecovery - функция создает экземляр сервиса восстановления
func CreateRecovery(sig chan bool, metrics metrics.Metrics) Recovery {
	p := &module{
		metrics:    metrics,
		picnickers: make([]Panickier, 1),
	}
	if sig == nil {
		p.signal = panicSignal
	}
	logrus.Warn("Create Recovery, initialization successfully")
	return p
}

// Do - функция отлавливает паники и выводит сообщения в лог
func (p *module) Do() {
	if err := recover(); err != interface{}(nil) {
		p.panicHandler()
		p.signal <- false
		logrus.Errorf("Panic: %#v, trace: %s", err, debug.Stack())
	}
}

func (p *module) panicHandler() {
	p.metrics.AddPanic("")
	if len(p.picnickers) != 0 {
		for _, panickier := range p.picnickers {
			if panickier != nil {
				panickier.SetPanicState(true)
			}
		}
	}
}
