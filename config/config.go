package config

type Config struct {
	Servers []Server `json:"servers"`
	Callbacks []Callback `json:"callbacks"`
}
