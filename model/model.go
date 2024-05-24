package model

type Report struct {
	Report_id    int    `json:"report_id"`
	Hero_id      int    `json:"hero_id"`
	Villain_id   int    `json:"villain_id"`
	Description  string `json:"description"`
	IncidentTime string `json:"incident_time"`
}
