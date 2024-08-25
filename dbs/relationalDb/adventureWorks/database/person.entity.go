package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Person struct {
	Id             int `gorm:"column:BusinessEntityID"`
	NameStyle      bool
	Title          sql.NullString
	FirstName      string
	MiddleName     sql.NullString
	LastName       string
	Suffix         sql.NullString
	EmailPromotion int
	RowGuid        uuid.UUID `gorm:"type:uuid;column:rowguid;not null"`
	ModifiedDate   time.Time
	PersonTypeId   int
	BusinessEntity BusinessEntity `gorm:"foreignKey:Id;references:Id"`
	PersonType     PersonType     `gorm:"foreignKey:PersonTypeId;references:PersonTypeId"`
}

type Persons []*Person

func (u *Person) TableName() string {
	return "Person.Person"
}

type PersonType struct {
	Id                    uint      `gorm:"column:PersonTypeId;primaryKey"`
	PersonTypeGuid        uuid.UUID `gorm:"type:uuid"`
	PersonTypeCode        string
	PersonTypeName        string
	PersonTypeDescription string
	CreatedBy             uint
	CreatedOn             time.Time
	ModifiedBy            uint
	ModifiedOn            time.Time
}

func (u *PersonType) TableName() string {
	return "Person.PersonType"
}

type PersonTypes []*PersonType
