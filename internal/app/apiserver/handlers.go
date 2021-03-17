package apiserver

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/mfaulther/avito-tradex-test-task/internal/app/model"
	"net/http"
)

func (s *APIServer) getStats(w http.ResponseWriter, r *http.Request) {

	sort := r.URL.Query().Get("sort")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	err := validation.Validate(from, validation.Required, validation.Date("2006-01-02"))
	if err != nil {
		s.ErrorHandler(w, 400, err.Error())
		return
	}

	err = validation.Validate(to, validation.Required, validation.Date("2006-01-02"))
	if err != nil {
		s.ErrorHandler(w, 400, err.Error())
		return
	}

	res := s.repository.GetStatistics(from, to, sort)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)

}

func (s *APIServer) addStats(w http.ResponseWriter, r *http.Request) {

	var newStats model.Statistics

	e := json.NewDecoder(r.Body).Decode(&newStats)

	if e != nil {
		s.ErrorHandler(w, 400, "Invalidate data")
		return
	}

	e = newStats.Validate()

	if e != nil {
		s.ErrorHandler(w, 400, e.Error())
		return
	}

	s.repository.AddStatistics(&newStats)

	w.WriteHeader(201)
	resp := map[string]interface{}{"status": 201, "answer": "statistics successfully added !"}

	json.NewEncoder(w).Encode(&resp)

}

func (s *APIServer) delStats(w http.ResponseWriter, r *http.Request) {

	s.repository.DeleteStatistics()
	resp := map[string]interface{}{"status": 200, "answer": "all statistics has been successfully deleted"}
	json.NewEncoder(w).Encode(&resp)

}

func (s *APIServer) ErrorHandler(w http.ResponseWriter, statusCode int, message string) {

	w.WriteHeader(statusCode)
	resp := map[string]interface{}{"status": 400, "error": message}
	json.NewEncoder(w).Encode(&resp)

}
