package main

import (
	"database/sql"
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

	fmt.Println()

	pr := database.NewPersonRepository(database.DatabaseContext)
	michaelScott, err := pr.GetPersonById(1000003)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Michael Scott RowId: %s <<>> Business Entity RowId: %s \n ",
			michaelScott.RowGuid.String(),
			michaelScott.BusinessEntity.RowGuid.String())
	}

	//personTypes := make([]*database.PersonType, 0)
	//database.DatabaseContext.Find(&personTypes)

	persons, err := pr.GetPersons()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The total person number found is: %v \n", len(persons))
	}

	newPerson := &database.Person{
		Title:        sql.NullString{String: "Mr.", Valid: true},
		FirstName:    "A First Name",
		LastName:     "LName",
		PersonTypeId: 4,
	}

	_, err = pr.CreatePerson(newPerson)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("A new person was created with an id of %v \n", newPerson.Id)
	}

	rows, err := pr.DeletePerson(newPerson.Id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("A new person was delete with an id of %v.... Confirmed rows deleted: %v \n", newPerson.Id, rows)
	}

	fmt.Println("--> Exiting Application")
}
