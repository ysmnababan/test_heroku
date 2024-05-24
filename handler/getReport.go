package handler

import (
	"database/sql"
	"encoding/json"
	"heroku/model"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HandlerDB struct {
	*sql.DB
}

func (h *HandlerDB) GetReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var reports []model.Report

	query := "SELECT report_id, hero_id, villain_id, description, incidentTime FROM reports"

	rows, err := h.Query(query)
	if err != nil {
		http.Error(w, "Error while getting query", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var r model.Report
		err = rows.Scan(&r.Report_id, &r.Hero_id, &r.Villain_id, &r.Description, &r.IncidentTime)
		if err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		reports = append(reports, r)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
