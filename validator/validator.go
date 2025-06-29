package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type Validator struct {
	errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{
		errors: make(map[string]string),
	}
}

func (v *Validator) Validator(data interface{}) bool {

	v.errors = make(map[string]string)
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return false
	}
	typeOfData := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typeOfData.Field(i)
		tag := fieldType.Tag.Get("validate")

		if tag != "" {
			fieldname := fieldType.Name
			if !v.validate(field, tag, fieldname) {
				v.errors[fieldname] = fmt.Sprintf("validation faield for %s on rule %s", fieldname, tag)
			}
		}

	}
	return len(v.errors) == 0
}

func (v *Validator) validate(field reflect.Value, tag string, fieldname string) bool {

	rules := strings.Split(tag, ",")

	for _, rule := range rules {
		if !v.applyRule(field, rule) {
			v.errors[fieldname] = fmt.Sprintf("validation faield for %s on rule %s", fieldname, rule)
			return false
		}
	}
	return true
}

func (v *Validator) applyRule(field reflect.Value, rule string) bool {

	parts := strings.Split(rule, "=")

	switch parts[0] {
	case "required":
		return !isEmpty(field)
	case "email":
		return isvalidEmail(field.String())
	case "mobile":
		return isvalidMobile(field.String())
	}

	return true
}

func isEmpty(field reflect.Value) bool {
	return field.Interface() == reflect.Zero(field.Type()).Interface()
}

func isvalidEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
func isvalidMobile(mobile string) bool {
	var emailRegex = regexp.MustCompile(`^(\+98|0)?9\d{9}$`)
	return emailRegex.MatchString(mobile)
}
func (v *Validator) GetErrors() map[string]string {
	return v.errors
}
