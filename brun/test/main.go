package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"red-envelope/infra"
	_ "red-envelope"
)

func main()  {
	filePath := kvs.GetCurrentFilePath("config.ini",1)
	//可以加载多个配置文件
	fileConfig := ini.NewIniFileCompositeConfigSource(filePath)
	app := infra.New(fileConfig)
	app.Start()
	ch := make(chan int)
	<- ch
}

