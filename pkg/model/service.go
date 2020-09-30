package model

import (
	_ "fmt"
	"gin_gateway/pkg/utils"
)

type Service struct {
	BaseModel

	Name 		string `gorm:"not null;size:32"`
	Url 		string `gorm:"not null;size:255"`
	Desc 		string `gorm:"not null;size:255"`
	Type 		string `gorm:"not null;size:64"`
	IsDefault 	bool   `gorm:"not null;default:false"`
	Creator   	string `gorm:"not null;size:32"`
}

type ProjectService struct {
	BaseModel

	ProjectID 		string 		`gorm:"not null;size:32"`
	ServiceID 		string 		`gorm:"not null;size:32"`
	ServiceType 	string 		`gorm:"not null;size:64"`
	Config 			utils.JSON 	`gorm:"not null;default:{};type:json"`
}