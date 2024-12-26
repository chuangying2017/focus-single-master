// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PCate is the golang structure of table p_cate for DAO operations like Where/Data.
type PCate struct {
	g.Meta    `orm:"table:p_cate, do:true"`
	Id        interface{} // 主键
	Pid       interface{} // 用于关联上级分类0为顶级
	Status    interface{} // 1启动0禁用，默认不启用
	Name      interface{} // 分类名称
	UpdatedId interface{} // 操作人员
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
