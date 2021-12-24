package config

type System struct {
	Mysql Mysql `map:"mysql" json:"mysql" yaml:"mysql"`
	Server  Server  `map:"server" json:"server" yaml:"server"`
}