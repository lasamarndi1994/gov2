package helper

import (
	"math/rand"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func GenerateUserID() string {
	return strconv.Itoa(100000 + rand.Intn(900000)) // generates 6-digit number (100000–999999)
}
