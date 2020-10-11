package main

import (
	"log"

	versionControlMigration "gormigrate-demo/version"
	initMigration "gormigrate-demo/version/v0"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var migrationOptions = gormigrate.Options{
	// UseTransaction:            true, // 使用 PostgreSQL 事务机制。
	// ValidateUnknownMigrations: true,  // 使用非法 migrations 记录校验。
}

func initMigrate(db *gorm.DB) {
	m := gormigrate.New(db, &migrationOptions, []*gormigrate.Migration{})

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(initMigration.ModelSchemaList...)
		if err != nil {
			return err
		}
		return nil
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Init Migration successfully")

}

func upgradeLatestMigrate(db *gorm.DB) {
	m := gormigrate.New(db, &migrationOptions, versionControlMigration.ModelSchemaList)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Upgrade Latest Migration successfully")
}

func main() {
	postgresDSN := "host=172.18.22.204 port=5432 user=test dbname=test_db password=1qaz@WSX sslmode=disable"
	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// db.Debug()
	log.Printf("Current Database: %s", db.Migrator().CurrentDatabase())

	// Case1. 初始化 v0 版本，如果已经初始化则不会再次执行 DDL 操作。
	initMigrate(db)

	// Case2. 迁移至最新版本，如果已经初始化则不会再次执行 DDL 操作。
	upgradeLatestMigrate(db)
}
