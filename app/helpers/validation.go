package helpers

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate

func ValidationForDTO(data interface{}) (map[string]string, error) {
	validate = validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(data)
	if err != nil {
		var errors = make(map[string]string)
		for _, erx := range err.(validator.ValidationErrors) {
			errors[erx.StructField()] = erx.Translate(trans)
		}

		return errors, err
	}
	return nil, nil

}
