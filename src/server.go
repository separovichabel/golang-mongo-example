package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Server struct {
	config       *Config
	baseCService *BaseCService
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/getAll", s.getAll)
	mux.HandleFunc("/get/", s.get)

	return mux
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.ServerPort,
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}

func (s *Server) getAll(w http.ResponseWriter, r *http.Request) {
	events := s.baseCService.FindAll()
	bytes, _ := json.Marshal(events)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	cpf := strings.TrimPrefix(r.URL.Path, "/get/")
	events := s.baseCService.getEvents(cpf)
	bytes, _ := json.Marshal(events)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (s *Server) setEvent(w http.ResponseWriter, r *http.Request) {
	cpf := strings.TrimPrefix(r.URL.Path, "/provisions/")
	events := s.baseCService.getEvents(cpf)
	bytes, _ := json.Marshal(events)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(config *Config, service *BaseCService) *Server {
	return &Server{
		config:       config,
		baseCService: service,
	}
}
