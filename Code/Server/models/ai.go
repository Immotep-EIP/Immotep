package models

type SummarizeRequest struct {
	Type     string   `binding:"required,oneof=room furniture"       json:"type"`
	Id       string   `binding:"required"                            json:"id"`
	Pictures []string `binding:"required,min=1,dive,required,base64" json:"pictures"`
}

type SummarizeResponse struct {
	State       string `json:"state"`
	Cleanliness string `json:"cleanliness"`
	Note        string `json:"note"`
}

type CompareRequest struct {
	Type     string   `binding:"required,oneof=room furniture"       json:"type"`
	Id       string   `binding:"required"                            json:"id"`
	Pictures []string `binding:"required,min=1,dive,required,base64" json:"pictures"`
}

type CompareResponse struct {
	State       string `json:"state"`
	Cleanliness string `json:"cleanliness"`
	Note        string `json:"note"`
}
