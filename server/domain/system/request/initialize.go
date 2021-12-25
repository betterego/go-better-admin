package request

type DBConfig struct {
	Host     string `json:"host" form:"host"`  // 服务器地址
	Port     string `json:"port" form:"port"`  // 数据库连接端口
	UserName string `json:"userName" form:"userName"` // 数据库用户名
	Password string `json:"password" form:"password" binding:"required"`  // 数据库密码
	DBName   string `json:"dbName " form:"dbName" binding:"required"`   // 数据库名
}
