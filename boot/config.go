package boot

import (
	appconfig "YangCodeCamp/config"
	"YangCodeCamp/pkg/config"
)

func initConfig() {
	config.InitConfig("")
	appconfig.Initialize()
}
