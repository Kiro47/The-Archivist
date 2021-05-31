package database

import (
	"Archivist/config"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)
func databaseConnect() *gorm.DB {
	// https://gorm.io/docs/connecting_to_the_database.html
	switch strings.ToLower(config.Database.Type) {
	case "postgres":
		config.ZLog.Debug("Starting Database connection",
			zap.String("db_type", config.Database.Type),
			zap.String("user", config.Database.User),
			zap.String("password", "__USED__"),
			zap.String("protocol", config.Database.Protocol),
			zap.Int("port", config.Database.Port),
			zap.String("db_name", config.Database.DatabaseName),
			zap.Bool("ssl", config.Database.SSL),
		)
		// Connect
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			config.Database.Host, config.Database.User, config.Database.Password,
			config.Database.DatabaseName, config.Database.Port,
			strconv.FormatBool(config.Database.SSL),
			)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			config.ZLog.Panic("Failed to connect to Database.\n" + err.Error())
			os.Exit(-21)
		}
		return db
	case "mysql":
		config.ZLog.Debug("Starting Database connection",
			zap.String("db_type", config.Database.Type),
			zap.String("user", config.Database.User),
			zap.String("password", "__USED__"),
			zap.String("protocol", config.Database.Protocol),
			zap.Int("port", config.Database.Port),
			zap.String("db_name", config.Database.DatabaseName),
			zap.Bool("ssl", config.Database.SSL),
		)
		config.ZLog.Warn("Using Archivist with Mysql has not been tested")
	case "sqlite":
		config.ZLog.Debug("Starting Database connection",
			zap.String("db_type", config.Database.Type),
			zap.String("filepath", config.Database.Path),
		)
		db, err := gorm.Open(sqlite.Open(config.Database.Path), &gorm.Config{})
		if err != nil {
			config.ZLog.Panic("Failed to connect to Database.\n" + err.Error())
			os.Exit(-21)
		}
		return db
	default:
		config.ZLog.Panic("Unsupported database type, shutting down")
		os.Exit(-20)
	}
	// Technically unreachable
	return nil
}
func databaseStart() {
	db := databaseConnect()
	if db == nil {
		config.ZLog.Panic("Issue creating DB connection")
		os.Exit(-29)
	}
	// TODO: Migrations and init
}
