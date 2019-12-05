package app

import (
	"github.com/google/gops/agent"
	"study_gin_admin/internal/app/config"
)

func InitMonitor() error {
	if c:= config.Global().Monitor; c.Enable {
		err := agent.Listen(agent.Options{
			Addr:            c.Addr,
			ConfigDir:       c.ConfigDir,
			ShutdownCleanup: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
