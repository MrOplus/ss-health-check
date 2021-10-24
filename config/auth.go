package config

import (
	"encoding/base64"
	"fmt"
)

type BasicAuthentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth BasicAuthentication) ToBase64() string {
	if auth.Password == "" && auth.Username == "" {
		return ""
	}
	formatted := fmt.Sprintf("%s:%s",auth.Username,auth.Password)
	return base64.StdEncoding.EncodeToString([]byte(formatted))
}