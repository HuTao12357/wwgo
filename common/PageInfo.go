package common

import (
	"fmt"
	"gorm.io/gorm"
)

type PageInfo struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

func PageVO(pageNum int, pageSize int, dto string, db *gorm.DB) *PageInfo {
	var total int64
	if pageNum == 0 {
		pageNum = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	if err := db.Table(dto).Count(&total).Error; err != nil {
		fmt.Println(err)
	}
	return &PageInfo{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    int(total),
	}
}
