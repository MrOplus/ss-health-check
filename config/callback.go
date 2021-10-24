package config

type Callback struct {
	Type   string `json:"type"`
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Server string `json:"server,omitempty"`
	Auth  BasicAuthentication `json:"auth"`
	Url    string `json:"url,omitempty"`
	Method string `json:"method,omitempty"`
	Params string `json:"params,omitempty"`
}
