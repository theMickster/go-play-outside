package main

import (
	"fmt"
	"main/database"
	"os"
	"strconv"
)

func retrieveSqlConfig() database.Config {
	dbPort, err := strconv.Atoi(os.Getenv("AdventureWorksSqlServerPort"))
	if err != nil {
		message := fmt.Sprintf("The input sql port '%s' found in environment variables was not an integer", os.Getenv("AdventureWorksSqlServerPort"))
		panic(message)
	}
	return database.Config{
		DbServer:        os.Getenv("AdventureWorksSqlServerName"),
		DbName:          os.Getenv("AdventureWorksSqlDatabaseName"),
		DbPort:          dbPort,
		AuthMethod:      database.Windows,
		ApplicationName: "Adventures in Go",
	}
}

func CreateDbContext() {
	database.InitDatabase(retrieveSqlConfig())
}
