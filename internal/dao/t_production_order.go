// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"xlsx2mysql/internal/dao/internal"
)

// internalTProductionOrderDao is internal type for wrapping internal DAO implements.
type internalTProductionOrderDao = *internal.TProductionOrderDao

// tProductionOrderDao is the data access object for table t_production_order.
// You can define custom methods on it to extend its functionality as you wish.
type tProductionOrderDao struct {
	internalTProductionOrderDao
}

var (
	// TProductionOrder is globally public accessible object for table t_production_order operations.
	TProductionOrder = tProductionOrderDao{
		internal.NewTProductionOrderDao(),
	}
)

// Fill with you ideas below.
