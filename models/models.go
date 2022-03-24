package models

import "gorm.io/gorm"

type Prices struct {
	gorm.Model
	Unix int64 `json:"unix"`
	Symbol string `json:"symbol"`
	Open float64 `json:"open"`
	High float64 `json:"high"`
	Low float64 `json:"low"`
	Close float64 `json:"close"`
}

type Uploads struct {
	gorm.Model
	FileName string `json:"file_name" gorm:"index:idx_name,unique"`
	Finish bool `json:"finish"`
	ExpireTimeInSeconds int `json:"expire_time"`
}

type Chunks struct {
	gorm.Model
	UploadsID uint `json:"-"`
	Upload Uploads `json:"upload,omitempty" gorm:"foreignKey:UploadsID;references:ID"`
	PartNum int `json:"part_num"`
	Size int `json:"size"`
	Hash string `json:"hash"`
}
