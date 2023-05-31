package model

import (
	"supplyChainPrivacyComputeSystem/utils"

	"gorm.io/gorm"
)

type Result struct {
	ID      int    `json:"id"`
	UID     int    `json:"uid"`
	Res     string `json:"res"`
	Date    string `json:"date"`
	ResDesc string `json:"res_desc"`
}

func (r Result) TableName() string {
	return "result"
}

func CreateResult(result Result) *gorm.DB {
	return utils.DB.Create(&result)
}
