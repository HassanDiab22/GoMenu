package utils

import (
	"strconv"

	"gorm.io/gorm"
)

// Paginate applies pagination to any GORM model query
func Paginate(db *gorm.DB, model interface{}, pageStr, pageSizeStr string) (int64, error) {
	var total int64

	// Convert query params to integers, default to -1 (fetch all)
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	// If pagination params are missing or invalid, return all records
	if page <= 0 || pageSize <= 0 {
		result := db.Find(model)
		return result.RowsAffected, result.Error
	}

	// Get total count
	db.Model(model).Count(&total)

	// Apply pagination
	offset := (page - 1) * pageSize
	result := db.Limit(pageSize).Offset(offset).Find(model)

	return total, result.Error
}
