// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PCate is the golang structure for table p_cate.
type PCate struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键"`
	Pid       uint64      `json:"pid"       orm:"pid"        description:"用于关联上级分类0为顶级"`
	Status    uint        `json:"status"    orm:"status"     description:"1启动0禁用，默认不启用"`
	Name      string      `json:"name"      orm:"name"       description:"分类名称"`
	UpdatedId uint64      `json:"updatedId" orm:"updated_id" description:"操作人员"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
