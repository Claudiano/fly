package services

import (
	"fly-go/dtos"
	"fly-go/models"
	"fly-go/repositories"
)

var passagemRepository = repositories.PassagemRepository{}

type PassagemService struct {
}

func (PassagemService) CadastrarPassagem(passagemDto dtos.PassagemDto) (models.Passagem, error) {
	passagem, err := passagemRepository.Save(passagemDto)
	return passagem, err
}

func (PassagemService) CarregarPassagens() ([]models.Passagem, error) {
	passagens, err := passagemRepository.FindByAll()
	if err != nil {
		return passagens, err
	}
	return passagens, nil
}

func (PassagemService) CarregarPassagem(idPassagem uint64) (models.Passagem, error) {

	passagem, err := passagemRepository.FindById(idPassagem)

	if err != nil {
		return passagem, err
	}
	return passagem, nil
}

func (PassagemService) AtualizarPassagem(passagem models.Passagem) (models.Passagem, error) {
	result, err := passagemRepository.Update(passagem)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (PassagemService) ExcluirPassagem(passagem models.Passagem) (models.Passagem, error) {
	result, err := passagemRepository.Delete(passagem)
	if err != nil {
		return result, err
	}
	return result, nil

}
