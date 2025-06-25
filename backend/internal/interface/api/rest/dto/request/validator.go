package request

import (
	"backend/internal/domain/entities"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	utTrans "github.com/go-playground/validator/v10/translations/en"
)

func isValidUsername(fl validator.FieldLevel) bool {
	if entities.IsValidUsername(fl.Field().String()) {
		return true
	}
	return false
}

func NewValidate() (*validator.Validate, ut.Translator) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	english := en.New()
	trans, _ := ut.New(english, english).GetTranslator("en")
	utTrans.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" || name == "" {
			return ""
		}
		return name
	})

	registerCustomValidation(validate, trans, "safe_username",
		"The {0} must be composed of only alphanumeric, '-', or '_' character.",
		isValidUsername,
	)

	return validate, trans
}

func registerCustomValidation(validate *validator.Validate, trans ut.Translator, name, translation string, validationFunc validator.Func) {
	validate.RegisterValidation(name, validationFunc)
	validate.RegisterTranslation(name, trans, func(trans ut.Translator) error {
		return trans.Add(name, translation, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(name, fe.Field())
		return t
	})
}

var validate, _ = NewValidate()
