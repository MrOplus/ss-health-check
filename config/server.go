package config

type Server struct {
	Type          string `json:"type"`
	Address       string `json:"address"`
	Port          int    `json:"port"`
	Password      string `json:"password"`
	Encryption    string `json:"encryption"`
	Obfs          string `json:"obfs,omitempty"`
	ObfsParam     string `json:"obfs-param,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	ProtocolParam string `json:"protocol-param,omitempty"`
}