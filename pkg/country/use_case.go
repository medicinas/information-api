package country

import "github.com/medicinas/information-api/pkg/entity"

type CountryUseCase struct {
	repository Repository
}

func NewUseCase(r Repository) *CountryUseCase {
	return &CountryUseCase{
		repository: r,
	}
}

func (c *CountryUseCase) Find(id uint) (*entity.Country, error) {
	panic("implement me")
}

func (c *CountryUseCase) FindByCode(email string) (*entity.Country, error) {
	panic("implement me")
}

func (c *CountryUseCase) FindAll() ([]*entity.Country, error) {
	panic("implement me")
}

func (c *CountryUseCase) Update(country *entity.Country) error {
	panic("implement me")
}

func (c *CountryUseCase) Save(country *entity.Country) (uint, error) {
	panic("implement me")
}

func (c *CountryUseCase) Delete(id uint) error {
	panic("implement me")
}
