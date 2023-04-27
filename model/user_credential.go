package model

type UserCredential struct {
	BaseModel
	UserName string `gorm:"unique;size:50;not null" json:"userName"`
	Password string `gorm:"not null" json:"-"`
	IsActive bool   `gorm:"default:false" json:"isActive"`
}

func (UserCredential) TableName() string {
	return "mst_user"
}
