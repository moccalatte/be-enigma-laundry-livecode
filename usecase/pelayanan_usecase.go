package usecase

import (
	"fmt"

	"be-enigma-laundry-livecode/model/entity"
	"be-enigma-laundry-livecode/repository"
)

type PelayananUseCase interface {
	FindById(id string) (entity.Pelayanan, error)
}

type pelayananUseCase struct {
	repo repository.PelayananRepository
}

func (u *pelayananUseCase) FindById(id string) (entity.Pelayanan, error) {

	pelayanan, err := u.repo.Get(id)
	if err != nil {
		return entity.Pelayanan{}, fmt.Errorf("user with ID %s not found", id)
	}

	return pelayanan, nil
}

func NewPelayananUseCase(repo repository.PelayananRepository) PelayananUseCase {
	return &pelayananUseCase{repo: repo}
}
