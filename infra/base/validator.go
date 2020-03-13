package base

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
	"red-envelope/infra"
)

var validate *validator.Validate
var translator ut.Translator

type ValidatorStarter struct {
	infra.BaseStarter
}

func Validator() *validator.Validate {
	return validate
}

func Translator() ut.Translator {
	return translator
}

func (v *ValidatorStarter) Init(ctx infra.StarterContext) {
	validate := validator.New()
	//创建消息国际化通用翻译器
	cn := zh.New()        //定义中文翻译器
	uni := ut.New(cn, cn) //参数一缺省，参数二默认
	var found bool
	translator, found = uni.GetTranslator("zh") //返回默认，如果没有返回缺省，且是否找到默认
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			logrus.Error(err)
		}
	} else { //没有找到翻译器
		logrus.Error("not found translator")
	}
}

func ValidateStruct(s interface{}) (err error) {
	//验证
	err = Validator().Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logrus.Error(err)
		}

		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				logrus.Error(err.Translate(translator))
			}
		}
		return err
	}
	return nil
}