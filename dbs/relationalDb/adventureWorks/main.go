package main

import (
	"fmt"
	"main/database"
)

func main() {
	defer func() {
		dbInstance, _ := database.DatabaseContext.DB()
		_ = dbInstance.Close()
	}()

	fmt.Println("--> Welcome to the AdventureWorks App built using Go + Gorm + Microsoft Sql Server")

	CreateDbContext()

	fmt.Printf("The count of business entities is: %v \n", database.CountBusinessEntities())

	fmt.Println("--> Exiting Application")
}
