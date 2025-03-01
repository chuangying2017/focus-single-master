// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"focus-single/internal/dao/internal"
)

// internalPCateDao is an internal type for wrapping the internal DAO implementation.
type internalPCateDao = *internal.PCateDao

// pCateDao is the data access object for the table p_cate.
// You can define custom methods on it to extend its functionality as needed.
type pCateDao struct {
	internalPCateDao
}

var (
	// PCate is a globally accessible object for table p_cate operations.
	PCate = pCateDao{
		internal.NewPCateDao(),
	}
)

// Add your custom methods and functionality below.
