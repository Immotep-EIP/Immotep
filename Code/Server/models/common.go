package models

type ArchiveRequest struct {
	Archive bool `json:"archive"`
}

type IdResponse struct {
	ID string `json:"id"`
}
