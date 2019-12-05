package app

import (
	"context"
	"go.uber.org/dig"
	"os"
	"study_gin_admin/internal/app/config"
	"study_gin_admin/pkg/logger"
)

type options struct {
	ConfigFile 	string
	ModelFile 	string
	WWWDir 		string
	SwaggerDir 	string
	MenuFile   string
	Version 	string
}

// option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetModelFile 设定casbin 模型配置文件
func SetModelFile(s string) Option {
	return func(o *options) {
		o.ModelFile = s
	}
}

// SetWWWDir 设定静态站点目录
func SetWWWDir(s string) Option {
	return func(o *options) {
		o.WWWDir = s
	}
}

// SetSwaggerDir 设定swagger目录
func SetSwaggerDir(s string) Option {
	return func(o *options) {
		o.SwaggerDir = s
	}
}

// SetMenuFile 设定菜单数据文件
func SetMenuFile(s string) Option {
	return func(o *options) {
		o.MenuFile = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Init 应用初始化
func Init(ctx context.Context, opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	err := config.LoadGlobal(o.ConfigFile)
	handleError(err)

	cfg := config.Global()
	logger.Printf(ctx, "服务启动，运行模式：%s 版本号：%s 进程号：%d", cfg.RunMode,o.Version, os.Getpid())

	if v := o.ModelFile; v != "" {
		cfg.Casbin.Model = v
	}
	if v := o.WWWDir;v != "" {
		cfg.WWW = v
	}
	if v := o.SwaggerDir; v != "" {
		cfg.Swagger = v
	}
	if v:= o.MenuFile; v!= "" {
		cfg.Menu.Data = v
	}
	loggerCall, err := InitLogger()
	if err != nil {
		logger.Errorf(ctx, err.Error())
	}

	err = InitMonitor()
	if err != nil {
		logger.Errorf(ctx,err.Error())
	}

	// 初始化图形验证码
	InitCaptcha()


}

func BuildContainer() (*dig.Container,func()){
	container := dig.New()

	return container, func() {

	}
}






