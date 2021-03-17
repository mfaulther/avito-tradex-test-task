package apiserver

import (
	"encoding/json"
	"github.com/mfaulther/avito-tradex-test-task/internal/app/model"
	"io"
	"net/http"
)

func (s *APIServer) getStats(w http.ResponseWriter, r *http.Request) {

	res := s.repository.GetStatistics()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)

}

func (s *APIServer) addStats(w http.ResponseWriter, r *http.Request) {

	var newStats model.Statistics

	e := json.NewDecoder(r.Body).Decode(&newStats)

	if e != nil {
		s.badRequest(w, "Invalidate data")
		return
	}

	s.repository.AddStatistics(&newStats)

}

func (s *APIServer) badRequest(w http.ResponseWriter, message string) {

	w.WriteHeader(400)
	io.WriteString(w, message)

}
