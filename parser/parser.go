package parser

import (
	"errors"
	"ss-health-check/config"
	"ss-health-check/models"
	"strings"
)

func Parse(server config.Server) (models.Proxy,error) {
	if strings.ToLower(server.Type) == "ssr" {
		ssr := models.ShadowsocksR{
			Shadowsocks:    models.Shadowsocks{
				Address:    server.Address,
				Port:       server.Port,
				Password:   server.Password,
				Encryption: server.Encryption,
			},
			Obfs:           server.Obfs,
			ObfsParams:     server.ObfsParam,
			Protocol:       server.Protocol,
			ProtocolParams: server.ProtocolParam,
		}
		return ssr,nil

	}else if strings.ToLower(server.Type) == "ss" {
		ss := models.Shadowsocks{
			Address:    server.Address,
			Port:       server.Port,
			Password:   server.Password,
			Encryption: server.Encryption,
		}
		return ss,nil
	}
	return nil,errors.New("invalid server type")
}