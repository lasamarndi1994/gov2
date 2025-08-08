package models

import (
	"time"
)

type Asset struct {
	Id              uint      `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	AssetCode       string    `json:"asset_code"`
	AssetName       string    `json:"asset_name" gorm:"size:256"`
	ProductCategory string    `json:"product_category"  gorm:"size:256"`
	SerialNumber    string    `json:"serial_number"  gorm:"size:256"`
	ProductDetails  string    `json:"product_details"  gorm:"size:256"`
	ProductCost     uint32    `json:"product_cost"  gorm:"size:256"`
	AssetImage      string    `json:"asset_image"  gorm:"size:256"`
	AssignDate      time.Time `json:"asset_date"  gorm:"type:date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdateAt        time.Time `json:"updated_at"`
}
