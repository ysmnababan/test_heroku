package handler

import (
	"database/sql"
	"encoding/json"
	"heroku/model"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *HandlerDB) GetReportById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	ReportId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var rep model.Report
	err = h.QueryRow("SELECT hero_id, villain_id, description, incidentTime FROM reports WHERE report_id = ?", ReportId).
		Scan(&rep.Hero_id, &rep.Villain_id, &rep.Description, &rep.IncidentTime)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "invalid id", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	rep.Report_id = ReportId
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Type-Content", "application/json")
	json.NewEncoder(w).Encode(rep)
}
