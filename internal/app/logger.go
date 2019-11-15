package app

import (
	"errors"
	"os"
	"path/filepath"
	"study_gin_admin/internal/app/config"
	"study_gin_admin/pkg/logger"
	loggerhook "study_gin_admin/pkg/logger/hook"
	loggergormhook "study_gin_admin/pkg/logger/hook/gorm"
)

func InitLogger() (func(), error) {
	c := config.Global().Log
	logger.SetLevel(c.Level)
	logger.SetFormatter(c.Format)


	var file *os.File
	if c.Output != "" {
		switch c.Output {
		case "stdout":
			logger.SetOutput(os.Stdout)
		case "stderr":
			logger.SetOutput(os.Stderr)
		case "file":
			if name := c.OutputFile;name != "" {
				os.MkdirAll(filepath.Dir(name), 0777)
				f, err := os.OpenFile(name, os.O_APPEND| os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					return nil, err
				}
				logger.SetOutput(f)
				file = f
			}
		}
	}
	var hook *loggerhook.Hook
	if c.EnableHook {
		switch c.Hook {
		case "gorm":
			hc := config.Global().LogGormHook
			var dsn string
			switch hc.DBType {
			case "mysql":
				dsn = config.Global().MySQL.DSN()
			case "sqlite3":
				dsn = config.Global().Sqlite3.DSN()
			case "postgres":
				dsn = config.Global().Postgres.DSN()
			default:
				return nil, errors.New("unknown db")
			}
			h := loggerhook.New(loggergormhook.New(&loggergormhook.Config{
				DBType:       hc.DBType,
				DSN:          dsn,
				MaxLifetime:  hc.MaxLifetime,
				MaxOpenConns: hc.MaxOpenConns,
				MaxIdleConns: hc.MaxIdleConns,
				TableName:    hc.Table,
			}),
				loggerhook.SetMaxWorkers(c.HookMaxThread),
				loggerhook.SetMaxQueues(c.HookMaxBuffer),
			)
			logger.AddHook(h)
			hook = h
		}
	}
	return func() {
		if file != nil {
			file.Close()
		}
		if hook != nil {
			hook.Flush()
		}
	},nil

}