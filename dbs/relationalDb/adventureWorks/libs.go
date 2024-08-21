package main

import (
	"main/database"
	"os"
)

func retrieveSqlConfig() database.Config {
	return database.Config{
		DbServer:   os.Getenv("AdventureWorksSqlServerName"),
		DbName:     os.Getenv("AdventureWorksSqlDatabaseName"),
		DbPort:     1433,
		AuthMethod: database.Windows,
	}
}

func CreateDbContext() {
	database.InitDatabase(retrieveSqlConfig())
}
