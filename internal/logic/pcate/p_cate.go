package pcate

import (
	"context"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func init(){
	service.RegisterPCate(New())
}

type sPCate struct{}

func New() *sPCate {
	return &sPCate{}
}

func (s *sPCate) CreateCategory(ctx context.Context, _ *model.PCateAddInput) (err error) {
	// 这里需要创建一个商品分类的逻辑，这里开始调用dao层的方法
	return
}

func (s *sPCate) GetCategory(ctx context.Context) (res []*model.PCateTreeItem, err error) {
	// 只是查询商品分类的树形结构
	return
}
