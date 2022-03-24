package db

import (
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"uploadLargeFile/models"
	"uploadLargeFile/settings"
	"time"
)

var log = settings.GetLogger()

var db *gorm.DB

func init() {
	var err error
	logLevel := gormLogger.Silent
	if settings.DbDebug {
		logLevel = gormLogger.Info
	}
	db, err = gorm.Open(settings.DbDialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: gormLogger.Default.LogMode(logLevel),
	})

	if err != nil {
		log.WithField("DB_Name", settings.DbName).WithError(err).Error("Error when init db")
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.WithError(err).Error("Error when init db connection")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(2)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	migrate()
}

func GetDB() *gorm.DB {
	return db
}

func migrate() {
	if err := db.AutoMigrate(&models.Prices{}); err != nil {
		panic(err.Error())
	}
	if err := db.AutoMigrate(&models.Uploads{}); err != nil {
		panic(err.Error())
	}
	if err := db.AutoMigrate(&models.Chunks{}); err != nil {
		panic(err.Error())
	}
}
