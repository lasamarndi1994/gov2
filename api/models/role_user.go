package models

type RoleUser struct {
	RoleId int `gorm:"primaryKey"`
	UserId int `gorm:"primaryKey"`
	// CreatedAt time.Time
	// DeletedAt gorm.DeletedAt
}
