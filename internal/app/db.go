package app

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"myAccount/internal/utils"
)

// LoadDB 链接数据库
func LoadDB(config *Config) (*gorm.DB, error) {
	switch config.DB.Type {
	case "sqlite":
		return loadSqlite(config)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", config.DB.Type)
	}
}

func loadSqlite(config *Config) (*gorm.DB, error) {
	if !utils.IsExist(config.DB.File) {
		if err := utils.CreateFile(config.DB.File); err != nil {
			return nil, err
		}
	}
	return gorm.Open(sqlite.Open(config.DB.File), &gorm.Config{})
}
