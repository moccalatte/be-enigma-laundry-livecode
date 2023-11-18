package usecase

import (
	"fmt"

	"be-enigma-laundry-livecode/model/entity"
	"be-enigma-laundry-livecode/repository"
)

type PelangganUseCase interface {
	FindById(id string) (entity.Pelanggan, error)
}

type pelangganUseCase struct {
	repo repository.PelangganRepository
}

func (u *pelangganUseCase) FindById(id string) (entity.Pelanggan, error) {

	pelanggan, err := u.repo.Get(id)
	if err != nil {
		return entity.Pelanggan{}, fmt.Errorf("user with ID %s not found", id)
	}

	return pelanggan, nil
}

func NewPelangganUseCase(repo repository.PelangganRepository) PelangganUseCase {
	return &pelangganUseCase{repo: repo}
}