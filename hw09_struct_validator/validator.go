package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrLen       = errors.New("invalid length error")
	ErrMin       = errors.New("value min error")
	ErrMax       = errors.New("value max error")
	ErrIn        = errors.New("input error")
	ErrNotStruct = errors.New("not struct error")
	ErrWrongRule = errors.New("wrong rule error")
	ErrRegexp    = errors.New("regexp error")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	sb := strings.Builder{}
	for _, e := range v {
		sb.WriteString(fmt.Sprintf("Field: %s, Error: %v\n", e.Field, e.Err))
	}
	return sb.String()
}

func Validate(v interface{}) error {
	var vErrors = ValidationErrors{}
	e := reflect.ValueOf(v)
	if e.Kind() != reflect.Struct {
		return ErrNotStruct
	}
	for i := 0; i < e.NumField(); i++ {

		field := e.Type().Field(i)
		varName := field.Name
		varTag := field.Tag
		if varTag == "" {
			continue
		}
		varValue := e.Field(i).Interface()
		if e.Field(i).Kind() == reflect.Slice {
			for _, sliceVal := range varValue.([]string) {
				err := validateValue(varTag, sliceVal)
				if err != nil {
					if !isValidationError(&err) {
						return err
					}
					vErrors = append(vErrors, ValidationError{Field: varName, Err: err})
				}
			}
		} else {
			err := validateValue(varTag, varValue)
			if err != nil {
				if !isValidationError(&err) {
					return err
				}
				vErrors = append(vErrors, ValidationError{Field: varName, Err: err})
			}
		}
	}
	if len(vErrors) > 0 {
		return vErrors
	}
	return nil
}

func isValidationError(err *error) bool {
	if errors.Is(*err, ErrLen) ||
		errors.Is(*err, ErrMin) ||
		errors.Is(*err, ErrMax) ||
		errors.Is(*err, ErrIn) ||
		errors.Is(*err, ErrNotStruct) ||
		errors.Is(*err, ErrWrongRule) ||
		errors.Is(*err, ErrRegexp) {
		return true
	}
	return false
}

func validateValue(varTag reflect.StructTag, varValue interface{}) error {
	if varTagVal := varTag.Get("validate"); varTagVal != "" {
		valRules := strings.Split(varTagVal, "|")
		for _, rawRule := range valRules {
			valRule := strings.Split(rawRule, ":")
			if len(valRule) != 2 {
				return ErrWrongRule
			}
			rule := valRule[0]
			val := valRule[1]
			switch rule {
			case "len":
				intVal, err := strconv.Atoi(val)
				if err != nil {
					return err
				}
				if intVal != len(varValue.(string)) {
					return ErrLen
				}
			case "min":
				intVal, err := strconv.Atoi(val)
				if err != nil {
					return err
				}
				if intVal > varValue.(int) {
					return ErrMin
				}
			case "max":
				intVal, err := strconv.Atoi(val)
				if err != nil {
					return err
				}
				if intVal < varValue.(int) {
					return ErrMax
				}
			case "in":
				inStr := strings.Split(val, ",")
				if !contains(inStr, fmt.Sprintf("%v", varValue)) {
					return ErrIn
				}
			case "regexp":
				matched, err := regexp.MatchString(val, fmt.Sprintf("%v", varValue))
				if err != nil {
					return err
				}
				if !matched {
					return ErrRegexp
				}
			default:
			}
		}
	}
	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
