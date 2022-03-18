package server

type RequestError struct {
	Error string `json:"error"`
}

type PopulateRequestBody struct {
	Path string `json:"path"`
}
