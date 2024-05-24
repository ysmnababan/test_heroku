package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *HandlerDB) DeleteReport(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	ReportId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	result, err := h.Exec("DELETE FROM reports WHERE report_id = ?", ReportId)
	if err != nil {
		http.Error(w, "error while deleting", http.StatusInternalServerError)
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

	w.WriteHeader(http.StatusOK)
}
