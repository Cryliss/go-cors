package types

// Application interface
type Application interface {
	CreateOutputFile(domain string, results []*Test) error
}

type Test struct {
	Acao    string            `json:"access-control-allow-origins"`
	Acac    string            `json:"access-control-allow-credentials"`
	Headers map[string]string `json:"headers"`
	Method  string            `json:"method"`
	Origin  string            `json:"origin"`
	Test    string            `json:"test"`
	URL     string            `json:"url"`
}
