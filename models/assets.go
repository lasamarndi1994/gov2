package models

import (
	"encoding/json"
	"time"
)

type Asset struct {
	Id              uint `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	AssignedUser    uint32
	AssetCode       *string
	AssetName       string
	ProductCategory *string
	SerialNumber    *string
	ProductDetails  *string
	ProductCost     uint32
	AssetImage      *json.Encoder
	AssignDate      *time.Time
}
