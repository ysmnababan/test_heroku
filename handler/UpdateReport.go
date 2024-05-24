package handler

import (
	"encoding/json"
	"heroku/model"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *HandlerDB) UpdateReport(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	ReportId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var report model.Report
	err = json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	query := `UPDATE reports 
	SET hero_id=?,
	villain_id = ?,
	description = ?,
	incidentTime = ?
	WHERE report_id = ?`

	result, err := h.Exec(query, report.Hero_id, report.Villain_id, report.Description, report.IncidentTime, ReportId)
	if err != nil {
		http.Error(w, "error while updating", http.StatusInternalServerError)
		return
	}

	affectedRow, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "error while getting affected row", http.StatusInternalServerError)
		return
	}

	if affectedRow == 0 {
		http.Error(w, "report not found", http.StatusInternalServerError)
		return
	}

	report.Report_id = int(ReportId)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
