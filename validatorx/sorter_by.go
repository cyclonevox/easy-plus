package validatorx

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func checkSortBy(fl validator.FieldLevel) bool {
	params := strings.Split(fl.Param(), ",")
	sorters := strings.Split(fl.Field().String(), ",")

	for _, sorter := range sorters {
		sortBy := strings.Split(sorter, " ")
		if sortBy[0] == "" {
			return false
		}

		var exist bool
		for _, param := range params {
			if sortBy[0] == param {
				exist = true

				break
			}
		}

		if !exist {
			return false
		}
	}

	return true
}
