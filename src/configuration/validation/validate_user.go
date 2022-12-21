package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	"github.com/michaelcaxias/crud-golang/src/configuration/rest_error"
)

var (
	Validate = validator.New()
	translate universalTranslator.Translator
)

func init() {
	if validate, success := binding.Validator.Engine().(*validator.Validate); success { // casting
		english := en.New()
		unt := universalTranslator.New(english, english)
		translate, _ = unt.GetTranslator("en")
		enTranslation.RegisterDefaultTranslations(validate, translate)
	}
}

func ValidateUserError(validation_error error) *rest_error.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_error, &jsonErr) {
		return rest_error.ThrowBadRequestError("Campo do tipo inválido")
	} else if errors.As(validation_error, &jsonValidationError) {
		errorsCauses := []rest_error.Causes{}

		for _, err := range validation_error.(validator.ValidationErrors) {
			cause := rest_error.Causes{
				Message: err.Translate(translate),
				Field: err.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_error.ThrowRequestValidationError("Campos estão inválidos", errorsCauses)
	} else {
		return rest_error.ThrowBadRequestError("Erro ao tentar converter campos")
	}
}