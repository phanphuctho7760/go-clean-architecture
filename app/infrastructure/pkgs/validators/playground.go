package validators

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type playGroundValidator struct {
	v     *validator.Validate
	trans ut.Translator
}

func newPlayGroundValidator() ValidatorItf {
	v := validator.New()
	english := en.New()
	uni := ut.New(english, english)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	// _ = en_translations.RegisterDefaultTranslations(v, trans)

	// Get message by json tag name
	registerTagName(v)

	// override default error
	registerTranslateOverride(v, trans)

	return &playGroundValidator{
		v:     v,
		trans: trans,
	}
}

// ValidateStructOneErr validate a struct
func (receiver *playGroundValidator) ValidateStructOneErr(i interface{}) (err error) {
	var errStr string
	eValidations := receiver.v.Struct(i)

	if eValidations == nil {
		return
	}

	errList, _ := eValidations.(validator.ValidationErrors)
	for _, e := range errList {
		errStr += e.Translate(receiver.trans) + ","
	}

	if strings.HasSuffix(errStr, ",") {
		errStr = errStr[:len(errStr)-1]
	}

	return errors.New(errStr)
}

// ValidateStructMoreErr validate a struct
func (receiver *playGroundValidator) ValidateStructMoreErr(i interface{}) (errs []error) {
	eValidations := receiver.v.Struct(i)
	errList, _ := eValidations.(validator.ValidationErrors)
	for _, e := range errList {
		errs = append(errs, errors.New(e.Translate(receiver.trans)))
	}
	return
}

func registerTagName(v *validator.Validate) {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		numOfPartNeedToSplit := 2
		firstIndex := 0
		name := strings.SplitN(fld.Tag.Get("json"), ",", numOfPartNeedToSplit)[firstIndex]
		if name == "-" {
			return ""
		}

		return name
	})
}

func registerTranslateOverride(validate *validator.Validate, trans ut.Translator) {
	// Start override lowercase required
	_ = validate.RegisterTranslation(
		"required",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		},
	)
	// End override lowercase required

	// Start override lowercase message
	_ = validate.RegisterTranslation(
		"lowercase",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("lowercase", "{0} must be a lowercase!", true) // see universal-translator for details
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("lowercase", fe.Field())
			return t
		},
	)
	// End override lowercase message
}
