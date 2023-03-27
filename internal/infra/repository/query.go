package repository

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/AI1411/go-grpc-graphql/internal/util"
)

func AddWhereEq(query *gorm.DB, columnName string, value interface{}) *gorm.DB {
	if util.IsNilOrEmpty(value) {
		return query
	}
	return query.Where(fmt.Sprintf("%s = ?", columnName), value)
}

func AddWhereGte(query *gorm.DB, columnName string, value interface{}) *gorm.DB {
	if util.IsNilOrEmpty(value) {
		return query
	}
	return query.Where(fmt.Sprintf("%s >= ?", columnName), value)
}

func AddWhereLte(query *gorm.DB, columnName string, value interface{}) *gorm.DB {
	if util.IsNilOrEmpty(value) {
		return query
	}
	return query.Where(fmt.Sprintf("%s <= ?", columnName), value)
}

func AddWhereLike(query *gorm.DB, columnName string, value string) *gorm.DB {
	if util.IsNilOrEmpty(value) {
		return query
	}
	return query.Where(fmt.Sprintf("%s LIKE ?", columnName), "%"+value+"%")
}

func AddWhereIn(query *gorm.DB, columnName string, values interface{}, includeNull bool) *gorm.DB {
	if util.IsNilOrEmpty(values) {
		return query
	}
	if includeNull {
		return query.Where(fmt.Sprintf("(%s IN (?) OR %s IS NULL)", columnName, columnName), values)
	}
	return query.Where(fmt.Sprintf("%s IN (?)", columnName), values)
}

func AddOffset(query *gorm.DB, offset int32) *gorm.DB {
	if util.IsNilOrEmpty(offset) {
		return query
	}
	return query.Offset(int(offset))
}

func AddLimit(query *gorm.DB, limit int32) *gorm.DB {
	if util.IsNilOrEmpty(limit) {
		return query
	}
	return query.Limit(int(limit))
}
