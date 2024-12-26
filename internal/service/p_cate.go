package service

import (
	"context"
	"focus-single/internal/model"
)

// IPCate defines the interface for the p_cate service
type IPCate interface {
	CreateCategory(context.Context, *model.PCateAddInput) error
	GetCategory(context.Context) ([]*model.PCateTreeItem, error)
}

var localPCate IPCate

func PCate() IPCate {
	if localPCate == nil {
		panic("implement not found for interface IPCate, forgot register?")
	}
	return localPCate
}

func RegisterPCate(i IPCate) {
	localPCate = i
}

