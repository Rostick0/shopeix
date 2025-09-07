package validation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

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

// StructValidator — общая функция валидации DTO
// func StructValidator[T any](w http.ResponseWriter, r *http.Request, dst *T) bool {
// 	// Декодим JSON
// 	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
// 		http.Error(w, "invalid request body", http.StatusBadRequest)
// 		return false
// 	}

// 	// Валидируем
// 	if err := validate.Struct(dst); err != nil {
// 		validationErrors := err.(validator.ValidationErrors)
// 		// resp := make(map[string]string)

// 		// for _, fieldErr := range validationErrors {
// 		// 	resp[fieldErr.Field()] = fieldErr.Tag()
// 		// }

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusBadRequest)
// 		_ = json.NewEncoder(w).Encode(map[string]interface{}{
// 			"errors": validationErrors.Error(),
// 		})
// 		return false
// 	}

// 	return true
// }
