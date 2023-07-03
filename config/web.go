package config

import "YangCodeCamp/pkg/config"

func init() {
	config.Add("web", func() map[string]interface{} {
		return map[string]interface{}{
			// server port
			"port": config.Load("web.port", "8000"),
		}
	})
}
