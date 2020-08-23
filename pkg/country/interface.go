package country

import "github.com/medicinas/information-api/pkg/entity"

type Reader interface {
	Find(id uint) (*entity.Country, error)
	FindByCode(email string) (*entity.Country, error)
	FindAll() ([]*entity.Country, error)
}

type Writer interface {
	Update(country *entity.Country) error
	Save(country *entity.Country) (uint, error)
	Delete(id uint) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	Reader
	Writer
}
