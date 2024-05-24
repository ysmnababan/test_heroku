package handler

import (
	"encoding/json"
	"heroku/model"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *HandlerDB) CreateReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var report model.Report
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	result, err := h.Exec("INSERT INTO reports (hero_id, villain_id, description, incidentTime) VALUES (?,?,?,?)", report.Hero_id, report.Villain_id, report.Description, report.IncidentTime)
	if err != nil {
		http.Error(w, "error while updating", http.StatusInternalServerError)
		return
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "error while getting the id", http.StatusInternalServerError)
		return
	}

	report.Report_id = int(lastId)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
