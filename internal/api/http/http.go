package http

import (
	"context"
	"net/http"
	"time"

	v1 "github.com/PandaGoL/api-project/internal/api/http/v1"
	"github.com/PandaGoL/api-project/internal/services"
	"github.com/PandaGoL/api-project/pkg/options"
	"github.com/gorilla/mux"
)

// Server - основной объект HTTP API сервера
type Server struct {
	router        *mux.Router
	httpSrv       *http.Server
	shutdownState bool
}

// Init - функция инициализирует и возвращает HTTP API сервер
func Init(service services.UserService) *Server {
	s := &Server{
		router: mux.NewRouter(),
	}

	// конфигурация роутера версии 1
	v1.InitAPI(s.router, service)

	lr := s.router
	s.httpSrv = &http.Server{
		Handler: lr,
		Addr:    options.Get().APIAddr,
	}

	return s
}

// Serve - функция для запуска HTTP API сервера
func (s *Server) Serve() (err error) {
	err = s.httpSrv.ListenAndServe()
	return
}

// Stop - функция остановки работы сервера
func (s *Server) Stop() (err error) {
	s.setShutdownState(true)
	// ждем завершения или просто "гасим" сервер
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return s.httpSrv.Shutdown(ctx)
}

// setShutdownState - функция устанавливает флаг завершения приложения в указанное состояние
func (s *Server) setShutdownState(state bool) {
	s.shutdownState = state
}
