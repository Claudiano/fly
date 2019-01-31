package services

import (
	"fly/dtos"
	"fly/models"
	"fly/repositories"
)

var vooRepository = repositories.VooRepository{}

type VooService struct{}

func (VooService) CadastrarVoo(vooDto dtos.VooDto) (models.Voo, error) {
	voo, err := vooRepository.Save(vooDto)
	if err != nil {
		return voo, err
	}
	return voo, nil
}

func (VooService) CarregarVoos() ([]models.Voo, error) {
	voos, err := vooRepository.FindByAll()
	return voos, err
}

func (VooService) CarregarVoo(idVoo uint64) (models.Voo, error) {
	voo, err := vooRepository.FindById(idVoo)

	return voo, err
}

func (VooService) AtualizarVoo(voo models.Voo) error {

	err := vooRepository.Update(voo)
	return err
}

func (VooService) ExcluirVoo(voo models.Voo) error {
	err := vooRepository.Delete(voo)
	return err
}
