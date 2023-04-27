package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Vehicle struct {
	BrandID        string     `json:"brandId"`
	Brand          Brand      `json:"brand"`
	Model          string     `gorm:"varchar;size:30" json:"model"`
	ProductionYear int        `gorm:"size:4" json:"productionYear"`
	Color          string     `gorm:"varchar;size:30" json:"color"`
	IsAutomatic    bool       `json:"isAutomatic"`
	Stock          int        `gorm:"check:stock >= 0" json:"stock"`
	SalePrice      int        `gorm:"check:sale_price > 0" json:"salePrice"`
	Status         string     `gorm:"check:status IN ('baru', 'bekas')" json:"status"`
	Customers      []Customer `gorm:"many2many:customer_vehicles;" json:"customers,omitempty"`
	ImgPath        string     `json:"ImgPath,omitempty"`
	UrlPath        string     `json:"urlPath,omitempty"`
	BaseModel
}

func (v Vehicle) TableName() string {
	return "mst_vehicle"
}

func (v Vehicle) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.BrandID, validation.Required),
		validation.Field(&v.Model, validation.Required, validation.Length(3, 35)),
		validation.Field(&v.ProductionYear, validation.Required),
		validation.Field(&v.Color, validation.Required),
		validation.Field(&v.IsAutomatic, validation.In(true, false)),
		validation.Field(&v.Stock, validation.Required, validation.Min(0)),
		validation.Field(&v.SalePrice, validation.Required, validation.Min(1)),
		validation.Field(&v.Status, validation.In("baru", "bekas")),
	)
}
