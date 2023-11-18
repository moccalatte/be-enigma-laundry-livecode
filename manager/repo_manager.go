package manager

import (
	"be-enigma-laundry-livecode/database"
	"be-enigma-laundry-livecode/repository"
)

type RepoManager interface {
	PelangganRepo() repository.PelangganRepository
	PelayananRepo() repository.PelayananRepository
	TransaksiRepo() repository.TransaksiRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) PelangganRepo() repository.PelangganRepository {
	return repository.NewPelangganRepository(&database.DB{DB: r.infra.Conn()})
}

func (r *repoManager) PelayananRepo() repository.PelayananRepository {
	return repository.NewPelayananRepository(&database.DB{DB: r.infra.Conn()})
}

func (r *repoManager) TransaksiRepo() repository.TransaksiRepository {
	return repository.NewTransaksiRepository(&database.DB{DB: r.infra.Conn()})
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
