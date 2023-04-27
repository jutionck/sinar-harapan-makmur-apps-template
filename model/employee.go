package model

import "time"

type Employee struct {
	BaseModel
	FirstName        string `gorm:"size:30"`
	LastName         string `gorm:"size:30"`
	Address          string
	Email            string `gorm:"unique;size:30"`
	PhoneNumber      string `gorm:"unique;size:15"`
	Bod              time.Time
	Position         string
	Salary           int64 `gorm:"default:0"`
	ManagerID        *string
	Manager          *Employee `gorm:"foreignKey:ManagerID"`
	UserCredentialID string
	UserCredential   UserCredential `gorm:"foreignKey:UserCredentialID;unique"`
}

func (Employee) TableName() string {
	return "mst_employee"
}
