package database

import (
	"time"

	"github.com/google/uuid"
)

type BusinessEntity struct {
	Id           int       `gorm:"column:BusinessEntityID;autoIncrement"`
	RowGuid      uuid.UUID `gorm:"column:rowguid; type:uuid"`
	ModifiedDate time.Time
}

type BusinessEntities []*BusinessEntity

func (u *BusinessEntity) TableName() string {
	return "Person.BusinessEntity"
}

func CountBusinessEntities() int64 {
	var businessEntities []BusinessEntity

	result := DatabaseContext.Find(&businessEntities)
	return result.RowsAffected
}
