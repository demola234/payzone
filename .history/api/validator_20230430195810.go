package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/demola234/payzone/utils"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.isSuppoertedCurrency(currency)
	}
	return false
}
