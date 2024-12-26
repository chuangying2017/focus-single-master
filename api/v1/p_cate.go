package v1

import (
	"focus-single/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type PCateGetReq struct {
	g.Meta `path:"/cate/get" method:"get" summary:"获取分类列表" tags:"分类"`
}

type PCateGetRes struct {
	List   []*model.PCateTreeItem `json:"list" dc:"分类列表"`
}

type PCateAddReq struct {
	g.Meta `path:"/cate/add" method:"post" summary:"添加分类" tags:"分类"`
	Name string `json:"name" v:"required#请输入分类名称" dc:"分类名称"`
	Pid uint64 `json:"pid" v:"pid@integer|min:0#|请选择上级分类" dc:"上级分类"`
}

type PCateAddRes struct {
}