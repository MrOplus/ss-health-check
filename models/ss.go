package models

import "fmt"

type Shadowsocks struct {
	Address string
	Port int
	Password string
	Encryption string
}

func (shadowsocks Shadowsocks) ToMap() map[string]interface{} {
	conf := make(map[string]interface{})
	conf["type"] = "ss"
	conf["name"] = fmt.Sprintf("%s:%d",shadowsocks.Address,shadowsocks.Port)
	conf["server"] = shadowsocks.Address
	conf["port"] = shadowsocks.Port
	conf["password"] = shadowsocks.Password
	conf["cipher"] = shadowsocks.Encryption
	return conf
}