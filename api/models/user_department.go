package models

import "time"

type UserDepartment struct {
	Id           uint      `json:"id"`
	UserId       uint      `json:"user_id" gorm:"not null"`
	DepartmentId uint      `json:"department_id" gorm:"not null"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}
