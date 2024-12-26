// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PCateDao is the data access object for the table p_cate.
type PCateDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns PCateColumns // columns contains all the column names of Table for convenient usage.
}

// PCateColumns defines and stores column names for the table p_cate.
type PCateColumns struct {
	Id        string // 主键
	Pid       string // 用于关联上级分类0为顶级
	Status    string // 1启动0禁用，默认不启用
	Name      string // 分类名称
	UpdatedId string // 操作人员
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// pCateColumns holds the columns for the table p_cate.
var pCateColumns = PCateColumns{
	Id:        "id",
	Pid:       "pid",
	Status:    "status",
	Name:      "name",
	UpdatedId: "updated_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPCateDao creates and returns a new DAO object for table data access.
func NewPCateDao() *PCateDao {
	return &PCateDao{
		group:   "default",
		table:   "p_cate",
		columns: pCateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PCateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PCateDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PCateDao) Columns() PCateColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PCateDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PCateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PCateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
