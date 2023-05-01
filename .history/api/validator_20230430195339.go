package api

import "github.com/go-playground/validator/v10"

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
 if currency, ok := fieldLevel.Field().Interface().(string); ok {
	  if currency == "USD" || currency == "EUR" || currency == "GBP" {
   return true
  }
  return false
 }
}