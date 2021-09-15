package runner

func readFile(path string) string {
	return ""
}

func writeFile(path string, value string) {

}

type requestParams struct {
	Method  string            `json:"method"`
	Uri     string            `json:"uri"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

type response struct {
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
}

func request(params requestParams) response {
	return response{}
}
