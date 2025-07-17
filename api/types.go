package api

type Request struct {
	Type string `json:"type"`
}

type AllMidsRequest struct {
	Request
	Dex string `json:"dex,omitempty"`
}

type MetadataRequest struct {
	Request
	Dex string `json:"dex,omitempty"`
}

type ClearinghouseStateRequest struct {
	Request
	User string `json:"user"`
	Dex  string `json:"dex,omitempty"`
}

type OpenOrdersRequest struct {
	Request
	User string `json:"user"`
	Dex  string `json:"dex,omitempty"`
}
