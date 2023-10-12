// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-starter/internal/dao/internal"
)

// internalProvincesDao is internal type for wrapping internal DAO implements.
type internalProvincesDao = *internal.ProvincesDao

// provincesDao is the data access object for table provinces.
// You can define custom methods on it to extend its functionality as you wish.
type provincesDao struct {
	internalProvincesDao
}

var (
	// Provinces is globally public accessible object for table provinces operations.
	Provinces = provincesDao{
		internal.NewProvincesDao(),
	}
)

// Fill with you ideas below.
