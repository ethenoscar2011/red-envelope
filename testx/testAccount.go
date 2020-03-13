package testx

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"red-envelope/infra"
	"red-envelope/infra/base"
)

func init() {
	filePath := kvs.GetCurrentFilePath("../brun/test/config.ini",1)
	//可以加载多个配置文件
	fileConfig := ini.NewIniFileCompositeConfigSource(filePath)
	//注册starter
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDataBaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	app := infra.New(fileConfig)
	app.Start()
}
