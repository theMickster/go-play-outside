package database

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type AuthenticationMethod int

const (
	UserAndPassword AuthenticationMethod = iota + 1
	Windows
	AzureManagedIdentity
)

type Config struct {
	DbServer        string
	DbName          string
	DbUser          string
	DbPass          string
	ApplicationName string
	DbPort          int32
	AuthMethod      AuthenticationMethod
}

var DatabaseContext *gorm.DB

func InitDatabase(config Config) error {
	var err error
	var connectionString = fmt.Sprintf("Data Source=%s,%v;Database=%s;Application Name=%s;Encrypt=true;TrustServerCertificate=true;",
		config.DbServer, config.DbPort, config.DbName, config.ApplicationName)

	switch config.AuthMethod {
	case UserAndPassword:
		connectionString += fmt.Sprintf("User Id=%s;Password=%s;", config.DbUser, config.DbPass)
	case Windows:
		connectionString += "Trusted_Connection=True;"
	}

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	pool, _ := db.DB()
	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(100)
	pool.SetConnMaxLifetime(time.Hour)
	pool.SetConnMaxLifetime(time.Hour)

	DatabaseContext = db

	return nil
}
