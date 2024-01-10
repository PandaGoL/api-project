package http

import (
	"context"
	"net/http"
	"time"

	v1 "github.com/PandaGoL/api-project/internal/api/http/v1"
	"github.com/PandaGoL/api-project/internal/services"
	"github.com/PandaGoL/api-project/pkg/options"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	uuid "github.com/satori/go.uuid"
)

// Server - основной объект HTTP API сервера
type Server struct {
	router        *mux.Router
	httpSrv       *http.Server
	systemService services.SystemService
	panicState    bool
	shutdownState bool
}

// Init - функция инициализирует и возвращает HTTP API сервер
func Init(userService services.UserService, systemService services.SystemService) *Server {
	s := &Server{
		router:        mux.NewRouter(),
		systemService: systemService,
	}

	// сервисные "ручки"
	s.router.Handle("/metrics", promhttp.Handler())
	s.router.HandleFunc("/health", s.handlerHealth)

	// конфигурация роутера версии 1
	v1.InitAPI(s.router, userService, s.middlewareRequestID)

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

// SetPanicState - функция устанавливает флаг наличия паник в указанное состояние
func (s *Server) SetPanicState(state bool) {
	s.panicState = state
}

// setShutdownState - функция устанавливает флаг завершения приложения в указанное состояние
func (s *Server) setShutdownState(state bool) {
	s.shutdownState = state
}

// обогощает контекст идентификатором запроса
func (s *Server) middlewareRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.NewV4().String()
			r = r.WithContext(context.WithValue(r.Context(), "request_id", requestID))
			next.ServeHTTP(w, r)
		},
	)
}
