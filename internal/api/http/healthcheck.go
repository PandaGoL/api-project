package http

import (
	"net/http"
)

func (s *Server) handlerHealth(w http.ResponseWriter, r *http.Request) {

	if s.panicState { // если паника, то возвращаем наружу 523-ий код
		w.WriteHeader(523)
		w.Write([]byte("Panic"))
		return
	}
	if s.shutdownState { // если был получен сигнал на выход, то возвращаем 503-ий код
		w.WriteHeader(503)
		w.Write([]byte("Shutdown"))
		return
	}

	if err := s.systemService.BDCheck(); err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Database offline"))
		return
	}
	w.Write([]byte("DataBase connected"))
	w.WriteHeader(200)
}
