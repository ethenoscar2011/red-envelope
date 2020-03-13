package red_envelope

import (
	"red-envelope/infra"
	"red-envelope/infra/base"
	_ "red-envelope/core/accounts"
	_ "red-envelope/apis/web"
)

func init()  {
	//注册starter
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDataBaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.InitailizerStarter{})

}
