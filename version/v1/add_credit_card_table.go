package v1

import (
	v0ModelSchema "gormigrate-demo/version/v0"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// AddCreditCardTable Credit Card V1 Migration
var AddCreditCardTable = gormigrate.Migration{
	ID: "v1",
	Migrate: func(tx *gorm.DB) error {
		// it's a good pratice to copy the struct inside the function,
		// so side effects are prevented if the original struct changes during the time
		type CreditCard struct {
			ID     string
			Number string
			UserID uint
			User   v0ModelSchema.User
			Name   string
		}
		return tx.AutoMigrate(&CreditCard{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("credit_cards")
	},
}
