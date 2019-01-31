package services

import (
	"fly-go/dtos"
	"fly-go/models"
	"fly-go/repositories"
	"fmt"
)

// Repository responsavel pela conexão com o banco de dados
var passageiroRepository = repositories.PassageiroRepository{}

type PassageiroService struct {
}

func (PassageiroService) CadastrarPassageiro(passageiroDto dtos.PassageiroDto) (models.Passageiro, error) {
	passageiro, err := passageiroRepository.Save(passageiroDto)
	if err != nil {
		return passageiro, err
	}

	return passageiro, nil
}

func (PassageiroService) BuscarPassageiros() ([]models.Passageiro, error) {
	fmt.Println("buscar passageiros")
	passageiros, err := passageiroRepository.FindByAll()
	if err != nil {
		return passageiros, err
	}
	return passageiros, nil
}

func (PassageiroService) BuscarPassageiro(idPassageiro uint64) (models.Passageiro, error) {
	passageiro, err := passageiroRepository.FindById(idPassageiro)
	if err != nil {
		return passageiro, err
	}
	return passageiro, nil
}

func (PassageiroService) AtualizarPassageiro(passageiro models.Passageiro) error {
	err := passageiroRepository.Update(passageiro)
	return err
}

func (PassageiroService) ExcluirPassageiro(passageiro models.Passageiro) {
	passageiroRepository.Delete(passageiro)
}

func (PassageiroService) BuscarPassageiroLogin(credenciais dtos.CredenciaisDto) models.Passageiro {
	passageiro := passageiroRepository.FIndByPassageiro(credenciais)
	return passageiro
}
