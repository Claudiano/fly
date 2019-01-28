package services

import (
	"fly/models"
	"fly/repositories"
)

var vooRepository = repositories.VooRepository{}

type VooService struct{}

func (VooService) CadastrarVoo(voo models.Voo) {
	vooRepository.Save(voo)
}

func (VooService) CarregarVoos() []models.Voo {
	voos := vooRepository.FindByAll()
	return voos
}

func (VooService) CarregarVoo(idVoo uint64) models.Voo {
	passagem := vooRepository.FindById(idVoo)
	return passagem
}

func (VooService) AtualizarVoo(voo models.Voo) {
	vooRepository.Update(voo)
}

func (VooService) ExcluirVoo(voo models.Voo) {
	vooRepository.Delete(voo)
}
