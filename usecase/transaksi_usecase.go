package usecase

import (
	"fmt"

	"be-enigma-laundry-livecode/model/entity"
	"be-enigma-laundry-livecode/repository"
)

type TransaksiUseCase interface {
	FindById(id string) (entity.Transaksi, error)
}

type transaksiUseCase struct {
	repo repository.TransaksiRepository
}

func (u *transaksiUseCase) FindById(id string) (entity.Transaksi, error) {

	transaksi, err := u.repo.Get(id)
	if err != nil {
		return entity.Transaksi{}, fmt.Errorf("user with ID %s not found", id)
	}

	return transaksi, nil
}

func NewTransaksiUseCase(repo repository.TransaksiRepository) TransaksiUseCase {
	return &transaksiUseCase{repo: repo}
}
