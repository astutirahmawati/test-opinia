package validations

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type validation struct {
	v *validator.Validate
}

func NewValidation(v *validator.Validate) *validation {
	return &validation{
		v: v,
	}
}

func (v *validation) Validation(request interface{}) error {
	v.v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	v.v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	err := v.v.Struct(request)
	if err != nil {
		return err
	}

	return nil
}
