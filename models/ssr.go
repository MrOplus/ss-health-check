package models

import "fmt"

type ShadowsocksR struct {
	Shadowsocks
	Obfs string
	ObfsParams string
	Protocol string
	ProtocolParams string
}

func (shadowsocksr ShadowsocksR) ToMap() map[string]interface{} {
	conf := make(map[string]interface{})
	conf["type"] = "ssr"
	conf["name"] = fmt.Sprintf("%s:%d",shadowsocksr.Shadowsocks.Address,shadowsocksr.Shadowsocks.Port)
	conf["server"] = shadowsocksr.Address
	conf["port"] = shadowsocksr.Port
	conf["password"] = shadowsocksr.Password
	conf["cipher"] = shadowsocksr.Encryption
	conf["obfs"] = shadowsocksr.Obfs
	conf["obfs-param"] = shadowsocksr.ObfsParams
	conf["protocol"] = shadowsocksr.Protocol
	conf["protocol-param"] = shadowsocksr.ProtocolParams
	return conf
}