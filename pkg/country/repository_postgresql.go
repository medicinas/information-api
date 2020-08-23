package country

import (
	"github.com/jinzhu/gorm"
	"github.com/medicinas/information-api/pkg/entity"
	"log"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func (cr *RepositoryImpl) Find(id uint) (*entity.Country, error) {
	panic("implement me")
}

func (cr *RepositoryImpl) FindByCode(email string) (*entity.Country, error) {
	panic("implement me")
}

func (cr *RepositoryImpl) FindAll() ([]*entity.Country, error) {
	panic("implement me")
}

func (cr *RepositoryImpl) Update(country *entity.Country) error {
	panic("implement me")
}

func (cr *RepositoryImpl) Delete(id uint) error {
	panic("implement me")
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

func (cr *RepositoryImpl) Save(country *entity.Country) (uint, error) {
	err := cr.db.Create(country).Error
	if err != nil {
		log.Println("Error: Could not persist the country from the repository {}", err)
		return -1, err
	}
	return country.ID, nil
}
