package validation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

type ErrorsFormat struct {
	Errors map[string]string
}

func InitLang(validate *validator.Validate) ut.Translator {
	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	_ = entranslations.RegisterDefaultTranslations(validate, trans)

	return trans
}

func StructValidator[T any](dst *T) map[string]string {
	var validate = validator.New()

	trans := InitLang(validate)

	err := validate.Struct(dst)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = e.Translate(trans) // красивое сообщение на англ
	}
	return errors
}
