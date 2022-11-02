package validateService

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
}

func TranslateError(data interface{}) (map[string]string) {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}
	errs := make(map[string]string)
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs[e.Field()] = translatedErr.Error()
	}
	return errs
}
