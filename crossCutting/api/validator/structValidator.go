package validator

import (
	"reflect"
	"sync"

	"github.com/go-playground/validator/v10"
)

type structValidator struct {
	//nolint:unused
	once     sync.Once
	validate *validator.Validate
}

func NewStructValidator() *structValidator {
	v := validator.New()
	v.SetTagName("binding")
	return &structValidator{
		validate: v,
	}
}

func (v *structValidator) ValidateStruct(obj interface{}) error {
	if v.kindOfData(obj) == reflect.Struct {
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}
func (v *structValidator) Engine() interface{} {
	return v.validate
}

func (v *structValidator) kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
