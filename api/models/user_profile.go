package models

type Profile struct {
	Id                     uint `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	Gender                 *string
	Address                *string
	PresentAdress          *string
	City                   *string
	State                  *string
	PrimaryContactNumber   *uint32
	SecondaryContactNumber *uint32
	PassportNo             *string
	PassportExpDate        *string
	Nationality            *string
	Religion               *string
	MaritalStatus          *string
	EmploymentOfSpouse     *string
	NoOfChildren           *int
}
