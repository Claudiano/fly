package services

import (
	"fly/dtos"
	"fly/models"
	"fly/repositories"
)

var passagemRepository = repositories.PassagemRepository{}

type PassagemService struct {
}

func (PassagemService) CadastrarPassagem(passagemDto dtos.PassagemDto) {
	passagemRepository.Save(passagemDto)
}

func (PassagemService) CarregarPassagens() []models.Passagem {
	passagens := passagemRepository.FindByAll()
	return passagens
}

func (PassagemService) CarregarPassagem(idPassagem uint64) models.Passagem {
	passagem := passagemRepository.FindById(idPassagem)
	return passagem
}

func (PassagemService) AtualizarPassagem(passagem models.Passagem) {
	passagemRepository.Update(passagem)
}

func (PassagemService) ExcluirPassagem(passagem models.Passagem) {
	passagemRepository.Delete(passagem)
}
