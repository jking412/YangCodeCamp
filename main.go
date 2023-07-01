package main

import (
	"YangCodeCamp/db"
	"YangCodeCamp/pkg/config"
	"YangCodeCamp/web"
)

func main() {
	config.InitConfig("config")
	db.Init()
	web.Init()
}
