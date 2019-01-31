package repositories

import (
	"fly/dtos"
	"fly/models"
	"fmt"
)

type PassageiroRepository struct {
}

func (PassageiroRepository) Save(passageiroDto dtos.PassageiroDto) (models.Passageiro, error) {
	var passageiro = models.Passageiro{}

	// criando passageiro
	//passageiro.IdPassageiro = 0
	passageiro.Nome = passageiroDto.Nome
	passageiro.Email = passageiroDto.Email
	passageiro.Cpf = passageiroDto.Cpf
	passageiro.Senha = passageiroDto.Senha

	db := connectar()

	err := db.Create(&passageiro).First(&passageiro).Error
	defer db.Close()

	fmt.Println(passageiro)
	if err != nil {
		fmt.Println(err)
		return passageiro, err
	}
	return passageiro, nil
}

func (PassageiroRepository) FindById(idPassageiro uint64) (models.Passageiro, error) {
	var passageiro models.Passageiro

	db := connectar()
	err := db.Where("IdPassageiro = ?", idPassageiro).Find(&passageiro).Error

	defer db.Close()

	if err != nil {
		return passageiro, err
	}

	return passageiro, nil

}

func (PassageiroRepository) FindByAll() ([]models.Passageiro, error) {
	fmt.Println("Consultando os passageiros na base")
	var passageiros []models.Passageiro
	db := connectar()
	err := db.Find(&passageiros).Error
	defer db.Close()

	if err != nil {
		return passageiros, err
	}

	return passageiros, nil
}

func (PassageiroRepository) Update(passageiro models.Passageiro) error {
	db := connectar()
	err := db.Save(&passageiro).Error
	defer db.Close()

	return err

}

func (PassageiroRepository) Delete(passageiro models.Passageiro) {
	db := connectar()

	db.Delete(&passageiro)

	defer db.Close()

}

func (PassageiroRepository) FIndByPassageiro(credenciais dtos.CredenciaisDto) models.Passageiro {
	var passageiro models.Passageiro
	db := connectar()

	db.Where("email =  ? and senha = ?", credenciais.Email, credenciais.Senha).First(&passageiro)

	defer db.Close()

	return passageiro

}
