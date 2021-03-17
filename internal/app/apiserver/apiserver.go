package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/mfaulther/avito-tradex-test-task/internal/app/repository"
	"net/http"
)

type APIServer struct {
	config     *Config
	router     *mux.Router
	repository *repository.StatRepository
}

func New(config *Config, repo *repository.StatRepository) (*APIServer, error) {

	return &APIServer{
		config:     config,
		router:     mux.NewRouter(),
		repository: repo,
	}, nil

}

func (s *APIServer) Start() error {

	s.router.HandleFunc("/stats", s.getStats).Methods("GET")
	s.router.HandleFunc("/stats", s.addStats).Methods("POST")
	s.router.HandleFunc("/stats", s.delStats).Methods("DELETE")
	e := http.ListenAndServe(s.config.BindAddr, s.router)
	if e != nil {
		return e
	}
	return nil
}
