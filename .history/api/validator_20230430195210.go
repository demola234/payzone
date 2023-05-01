package api

import "github.com/go-playground/validator/v10"

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {