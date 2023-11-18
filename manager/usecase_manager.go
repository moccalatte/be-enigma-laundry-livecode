package manager

import "be-enigma-laundry-livecode/usecase"

type UseCaseManager interface {
	PelangganUseCase() usecase.PelangganUseCase
	PelayananUseCase() usecase.PelayananUseCase
	TransaksiUseCase() usecase.TransaksiUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) PelangganUseCase() usecase.PelangganUseCase {
	return usecase.NewPelangganUseCase(u.repo.PelangganRepo())
}

func (u *useCaseManager) PelayananUseCase() usecase.PelayananUseCase {
	return usecase.NewPelayananUseCase(u.repo.PelayananRepo())
}

func (u *useCaseManager) TransaksiUseCase() usecase.TransaksiUseCase {
	return usecase.NewTransaksiUseCase(u.repo.TransaksiRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
