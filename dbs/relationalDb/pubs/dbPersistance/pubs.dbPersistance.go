package dbPersistance

import (
	"context"
	"database/sql"
	"fmt"
)

type Database struct {
	PubsDatabase *sql.DB
}

var pubsDbContext = context.Background()

func (db Database) IsAlive() error {
	err := db.PubsDatabase.PingContext(pubsDbContext)
	if err != nil {
		return err
	}

	sql := "SELECT COUNT(*) FROM authors;"
	rows, err := db.PubsDatabase.QueryContext(pubsDbContext, sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return err
		}
	}

	return nil
}

func (db Database) GetAuthors() ([]Author, error) {
	const errorText = "error retrieving authors. %v"
	var authors []Author

	sql := `SELECT a.author_id AS Id
,a.author_code AS AuthorCode
,a.last_name AS LastName
,a.first_name AS FirstName
,a.phone_number AS PhoneNumber
,COALESCE(a.[address], '') AS [Address]
,COALESCE(a.city, '') AS [City]
,COALESCE(a.[state], '') AS [State]
,COALESCE(a.zip_code, '') AS [ZipCode]
,a.[contract] AS [HasContract]
FROM dbo.authors a;`

	rows, err := db.PubsDatabase.QueryContext(pubsDbContext, sql)
	if err != nil {
		return nil, fmt.Errorf(errorText, err)
	}
	defer rows.Close()

	for rows.Next() {
		var a Author
		err := rows.Scan(&a.Id, &a.AuthorCode, &a.LastName, &a.FirstName, &a.PhoneNumber, &a.Address, &a.City, &a.State, &a.ZipCode, &a.HasContract)
		if err != nil {
			return nil, fmt.Errorf(errorText, err)
		}
		authors = append(authors, a)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(errorText, err)
	}

	return authors, nil
}
