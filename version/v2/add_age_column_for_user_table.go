package v2

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// AddAgeColumnForUserTable User v2 Migration
var AddAgeColumnForUserTable = gormigrate.Migration{
	ID: "v2",
	Migrate: func(tx *gorm.DB) error {
		// when table already exists, it just adds fields as columns
		type User struct {
			Age int
		}
		// NOTE(fanguiju): 这里可以选择使用 DDL 语句或 AutoMigrate 功能。
		return tx.AutoMigrate(&User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropColumn("users", "age")
	},
}
