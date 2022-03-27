package controller

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var translator ut.Translator

func RegisteValidator(local string) error {
	var ok bool
	register, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("获得gin默认validator出错")
	}
	//注册自定义的json tag的获得方法
	register.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	//注册 required
	//err := register.RegisterValidation("required1", requiredValidation)
	//if err != nil {
	//	return err
	//}
	zhT := zh.New()
	enT := en.New()
	uni := ut.New(zhT, enT, enT)
	translator, ok = uni.GetTranslator(local)
	if !ok {
		return errors.New("获取translator失败")
	}
	switch local {
	case "en":
		enTranslations.RegisterDefaultTranslations(register, translator)
	case "zh":
		zhTranslations.RegisterDefaultTranslations(register, translator)
	default:
		enTranslations.RegisterDefaultTranslations(register, translator)
	}
	return nil
}
func removeStructHeader(m map[string]string) map[string]string {
	newMap := make(map[string]string)
	for key, value := range m {
		value = value[strings.Index(value, ".")+1:]
		newMap[key] = value
	}
	return newMap
}
func requiredValidation(f validator.FieldLevel) bool {
	rs := f.Field()
	if !rs.IsValid() {
		return false
	}
	return true
}
func orderValidation(f validator.StructLevel) bool {
	order := f.Current().FieldByName("Order")
	if order.String() == "community_name" || order.String() == "create_date" || order.String() == "" {
		return true
	}
	return false
}
