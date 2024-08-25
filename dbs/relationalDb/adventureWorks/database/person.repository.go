package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository(input *gorm.DB) *PersonRepository {
	return &PersonRepository{db: input}
}

func (r *PersonRepository) GetPersonById(id uint) (*Person, error) {
	person := &Person{}
	if err := r.db.Model(&person).Preload("PersonType").Preload("BusinessEntity").First(person, id).Error; err != nil {
		return nil, err
	}
	return person, nil
}

func (r *PersonRepository) GetPersons() (Persons, error) {
	persons := make([]*Person, 0)
	if err := r.db.Limit(100).Find(&persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}

func (r *PersonRepository) CreatePerson(person *Person) (*Person, error) {
	businessEntity := &BusinessEntity{ModifiedDate: time.Now(), RowGuid: uuid.New()}
	if err := r.db.Create(businessEntity).Error; err != nil {
		return nil, err
	}

	person.Id = businessEntity.Id
	person.RowGuid = uuid.New()
	person.ModifiedDate = time.Now()

	if err := r.db.Create(person).Error; err != nil {
		return nil, err
	}
	return person, nil
}

func (r *PersonRepository) DeletePerson(id int) (int64, error) {
	result01 := r.db.Delete(&Person{}, id)
	result02 := r.db.Delete(&BusinessEntity{}, id)

	return result01.RowsAffected + result02.RowsAffected, result01.Error
}
