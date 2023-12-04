package common

import (
	"fmt"
	"gorm.io/gorm"
)

type PageInfo struct {
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
	Data     interface{} `json:"data" gorm:"-"`
}

func PageVO(pageNum int, pageSize int, dto interface{}, db *gorm.DB) *PageInfo {
	var total int64
	//var data interface{}
	if err := db.Model(dto).Count(&total).Error; err != nil {
		fmt.Println(err)
	}
	if err := db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(dto).Error; err != nil {
		fmt.Println(err)
	}
	return &PageInfo{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    int(total),
		Data:     dto,
	}
}
