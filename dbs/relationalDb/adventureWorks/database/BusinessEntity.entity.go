package database

import "time"

type BusinessEntity struct {
	BusinessEntityID uint   `gorm:"primaryKey"`
	RowGuid          string `gorm:"column:rowguid;not null"`
	ModifiedDate     time.Time
}

func (u *BusinessEntity) TableName() string {
	return "Person.BusinessEntity"
}

func CountBusinessEntities() int64 {
	var businessEntities []BusinessEntity

	result := DatabaseContext.Find(&businessEntities)
	return result.RowsAffected
}
