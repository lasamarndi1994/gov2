package models

type UserPremission struct {
	Id       uint `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	ModuleId uint32
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt
}
