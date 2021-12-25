package config

type System struct {
	Mysql Mysql `map:"mysql" json:"mysql" yaml:"mysql"`
	Zap Zap `map:"zap" json:"zap" yaml:"zap"`
	Server Server `map:"server" json:"server" yaml:"server"`
}