package api

type AllMidsRequest struct {
	Type string `json:"type"`
	Dex  string `json:"dex,omitempty"`
}

type MetadataRequest struct {
	Type string `json:"type"`
	Dex  string `json:"dex,omitempty"`
}
