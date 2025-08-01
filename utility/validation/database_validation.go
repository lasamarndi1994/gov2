package validation

import (
	"github.com/lasamarndi1994/gov2/internal/database"
)

func UniqueValidation(model interface{}, filed string, input string) bool {
	var count int64
	database.DB.Model(&model).Where(filed+" =?", input).Count(&count)
	if count == 0 {
		return true
	} else {
		return false
	}
}

// func NotFoundValidation(model interface{}, filed string, input string) error {
// 	if err := database.DB.Model(&model).Where(filed+" =?", input).Error; err != nil {
// 		return err
// 	}
// 	return
// }
