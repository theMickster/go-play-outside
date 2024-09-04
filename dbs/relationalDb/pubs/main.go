package main

import (
	"database/sql"
	"fmt"
	"os"
	"pubs/dbPersistance"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	var (
		serverName   = os.Getenv("PubsSqlServerName")
		databaseName = os.Getenv("PubsSqlDatabaseName")
	)

	connectionString := fmt.Sprintf("Data Source=%s;Database=%s;Application Name=Pubs RelationalDb via Go;Integrated Security=sspi;Encrypt=true;TrustServerCertificate=true;", serverName, databaseName)
	sqlObj, connectionError := sql.Open("mssql", connectionString)
	if connectionError != nil {
		fmt.Println(fmt.Errorf("oops. Something went wrong opening connection to SQL Server. %v", connectionError))
	}

	db := dbPersistance.Database{PubsDatabase: sqlObj}
	pubsHealthCheckErr := db.IsAlive()

	if pubsHealthCheckErr != nil {
		fmt.Println(fmt.Errorf("oops. Something went wrong with the Pubs health check. %v", pubsHealthCheckErr))
	}

	fmt.Println("--> Welcome to the Pubs Console App built using Go + Microsoft Sql Server")

	authors, err := db.GetAuthors()
	if err != nil {
		fmt.Println(fmt.Errorf("oops. Something went wrong with retrieving authors. %v", err))
	}
	fmt.Printf("The Authors located in the Pubs database are as follows \n\n %+v \n ", authors)

}
