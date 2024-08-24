package database

import (
	"fmt"
	"test-golang-muhamad-isro/helper"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection(DATABASE_URL string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(DATABASE_URL), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	helper.PanicIfError(err)

	sqlDB, err := db.DB()
	helper.PanicIfError(err)

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	defer fmt.Println("Successfully connect to database !")

	return db
}

func CloseConnection(db *gorm.DB) {

	dbInstance, err := db.DB()
	helper.PanicIfError(err)

	err = dbInstance.Close()
	helper.PanicIfError(err)

	fmt.Println("Connection Closed")
}
