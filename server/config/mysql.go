package config

type Mysql struct {
	Path         string `map:"path" json:"path" yaml:"path"`                             // 服务器地址:端口
	Config       string `map:"config" json:"config" yaml:"config"`                       // 高级配置
	Dbname       string `map:"db-name" json:"dbname" yaml:"dbname"`                     // 数据库名
	Username     string `map:"username" json:"username" yaml:"username"`                 // 数据库用户名
	Password     string `map:"password" json:"password" yaml:"password"`                 // 数据库密码
	MaxIdleConns int    `map:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `map:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `map:"log-mode" json:"logMode" yaml:"log-mode"`                  // 是否开启Gorm全局日志
	LogZap       bool   `map:"log-zap" json:"logZap" yaml:"log-zap"`                     // 是否通过zap写入日志文件
}

func (mysql *Mysql) Dsn() string {
	return mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Path + ")/" + mysql.Dbname + "?" + mysql.Config
}
