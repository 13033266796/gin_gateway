package model

import (
	_ "fmt"
	_ "gin_gateway/pkg/utils"
)

type Project struct {
	BaseModel

	Name 				string `gorm:"not null;size:64"`
	Status 				int64  `gorm:"not null"`
	ClientID 			string `gorm:"not null;size:64"`
	ClientSecret 		string `gorm:"not null;size:64"`
	Description 		bool   `gorm:""`
}
