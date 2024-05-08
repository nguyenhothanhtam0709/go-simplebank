package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/nguyenhothanhtam0709/go-simplebank/utils"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency)
	}

	return false
}
