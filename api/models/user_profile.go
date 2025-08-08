package models

type Profile struct {
	Id                     uint   `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	UserId                 uint   `json:"user_id"`
	Gender                 string `json:"gender" gorm:"size:256"`
	Address                string `json:"address" gorm:"size:256"`
	PresentAdress          string `json:"parment_address" gorm:"size:256"`
	City                   string `json:"city" gorm:"size:256"`
	State                  string `json:"state" gorm:"size:256"`
	PrimaryContactNumber   uint32 `json:"primary_conatct_number" gorm:"size:256"`
	SecondaryContactNumber uint32 `json:"secondary_conatct_number" gorm:"size:256"`
	PassportNo             string `json:"passport_no" gorm:"size:256"`
	PassportExpDate        string `json:"passport_exp_date" gorm:"size:256;type:date" `
	Nationality            string `json:"nationality" gorm:"size:256"`
	Religion               string `json:"religion" gorm:"size:256"`
	MaritalStatus          string `json:"marital_status" gorm:"size:256"`
	EmploymentOfSpouse     string `json:"employment_of_spouse" gorm:"size:256"`
	NoOfChildren           int    `json:"no_of_children" gorm:"size:256"`
	User                   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
