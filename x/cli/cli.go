package cli

// run 参数  配置文件

type Flag struct {
	Name     string
	FilePath string
	EnvVar   string
	Fn       func()
}
