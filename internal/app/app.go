package app

import (
	"context"
	"study_gin_admin/internal/app/config"
)

type options struct {
	ConfigFile 	string
	ModelFile 	string
	WWWDir 		string
	SwaggerDir 	string
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

func Init(ctx context.Context, opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	err := config.LoadGlobal(o.ConfigFile)
	handleError(err)

	cfg := config.Global()
}






