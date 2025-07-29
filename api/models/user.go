package models

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id           uint           `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Uuid         uuid.UUID      `json:"uuid" gorm:"size:256;not nul;unique"`
	UserId       string         `json:"emp_id" gorm:"unique;not null;uniqueIndex"`
	FirstName    string         `json:"first_name" gorm:"size:256"`
	LastName     string         `json:"last_name" gorm:"size:256;not null"`
	Email        string         `json:"email" gorm:"unique;not null;unique;unique_email"`
	MobileNumber int            `json:"mobile_number" gorm:"unique"`
	Password     string         `json:"-" gorm:"not null;size:256"`
	Status       string         `json:"status" gorm:"type:enum('Active','Pending','Approved','Deactive');default:'Pending';not null"`
	JoiningDate  time.Time      `json:"joining_date" gorm:"type:date;default:null"`
	DateofBirth  time.Time      `json:"dateof_birth" gorm:"type:date;default:null"`
	AboutMe      string         `json:"about_me"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Uuid = uuid.New() // NOT uuid.UUID{}
	u.UserId = strconv.Itoa(100000 + rand.Intn(900000))
	return
}
