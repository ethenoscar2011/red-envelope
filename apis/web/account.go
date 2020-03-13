package web

import (
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"red-envelope/infra"
	"red-envelope/infra/base"
	"red-envelope/services"
)

func init()  {
	infra.RegiserApi(new(AccountApi))
}

type AccountApi struct {

}

func (a *AccountApi)Init()  {

	groupRouter := base.Iris().Party("/v1/account")
	groupRouter.Post("/create", createHandler)
	groupRouter.Post("/transfer", transferHandler)

}

//创建账户 /v1/account/create
func createHandler(context iris.Context)  {
		logrus.Info("请求。。。")
		account :=services.AccountCreatedDTO{}
		err := context.ReadJSON(&account)
		if err != nil {
			context.JSON(base.Resp{-1,"创建账户失败",""})
			return
		}
		service := services.GetAccountService()
		dto,err := service.CreateAccount(account)
		if err != nil {
			context.JSON(base.Resp{-1,"创建账户失败",""})
			return
		}
		context.JSON(&dto)

}
//转账接口: /v1/account/transfer
func transferHandler(ctx iris.Context)   {
	account := services.AccountTransferDTO{}
	err := ctx.ReadJSON(&account)
	r := base.Resp{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Msg = err.Error()
		ctx.JSON(&r)
		return
	}
	//执行转账逻辑
	service := services.GetAccountService()
	status, err := service.Transfer(account)
	if err != nil {
		r.Code = base.ResCodeInnerServerError
		r.Msg = err.Error()
	}
	r.Data = status
	if status != services.TransferStatusSuccess {
		r.Code = base.ResCodeBizError
		r.Msg = err.Error()
	}
	ctx.JSON(&r)
}