package model

type Tb_casestudy struct {
	Id          int    `json:"id" gorm:"primary_key;not null"`
	Title       string `json:"title" gorm:"not null;default:null"`
	Description string `json:"description" gorm:"not null;default:null"`
	Imageuri    string `json:"imageuri" gorm:"not null;default:null"`
	Createddate string `json:"createddate" gorm:"not null;default:null"`
}
