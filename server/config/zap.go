package config

type Zap struct {
	Level         string `map:"level" json:"level" yaml:"level"`                           // 级别
	Format        string `map:"format" json:"format" yaml:"format"`                        // 输出
	Prefix        string `map:"prefix" json:"prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `map:"director" json:"director"  yaml:"director"`                 // 日志文件夹
	LinkName      string `map:"link-name" json:"linkName" yaml:"link-name"`                // 软链接名称
	ShowLine      bool   `map:"show-line" json:"showLine" yaml:"showLine"`                 // 显示行
	EncodeLevel   string `map:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `map:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `map:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}
