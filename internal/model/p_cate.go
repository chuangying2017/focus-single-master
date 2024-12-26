package model

import "github.com/gogf/gf/v2/os/gtime"

type PCateTreeItem struct {
	Id        uint64           `json:"id"`                 // 分类ID，自增主键
	Pid       uint64           `json:"pid"`                // 父级分类ID，用于层级管理
	Name      string           `json:"name"`               // 分类名称
	Children  []*PCateTreeItem `json:"children,omitempty"` // 子级数据项
	CreatedAt *gtime.Time      `json:"created_at"`         // 创建时间
	UpdatedAt *gtime.Time      `json:"updated_at"`         // 修改时间
}


type PCateAddInput struct {
	Name string
	Pid  uint64
}