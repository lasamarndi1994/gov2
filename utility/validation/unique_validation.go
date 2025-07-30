package validation

import (
	"fmt"

	"github.com/lasamarndi1994/gov2/internal/database"
)

func UniqueValidation(model interface{}, filed string, input string) bool {
	var count int64
	//mod = models
	database.DB.Model(&model).Where("name =?", input).Count(&count)
	fmt.Println(count)
	if count == 0 {
		return true
	} else {
		return false
	}
}
