package v3

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// AddEmailColumnForUserTable User v3 Migration
var AddEmailColumnForUserTable = gormigrate.Migration{
	ID: "v3",
	Migrate: func(tx *gorm.DB) error {
		// when table already exists, it just adds fields as columns
		type User struct {
			Email string
		}
		return tx.AutoMigrate(&User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropColumn("users", "email")
	},
}
