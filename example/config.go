package main

import (
	"fmt"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

func main()  {
	configPath := kvs.GetCurrentFilePath("config.ini",1)
	configSource := ini.NewIniFileConfigSource(configPath)
	port := configSource.GetIntDefault("app.server.port",10080)
	fmt.Println("port:",port)
}

