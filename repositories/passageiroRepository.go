package repositories

import (
	"fly/dtos"
	"fly/models"
	"fmt"
)

type PassageiroRepository struct {
}

func (PassageiroRepository) Save(passageiroDto dtos.PassageiroDto) {
	var passageiro = models.Passageiro{}

	// criando passageiro
	//passageiro.IdPassageiro = 0
	passageiro.Nome = passageiroDto.Nome
	passageiro.Email = passageiroDto.Email
	passageiro.Cpf = passageiroDto.Cpf
	passageiro.Senha = passageiroDto.Senha

	db := connectar()

	db.Create(&passageiro)
	defer db.Close()
}

func (PassageiroRepository) FindById(idPassageiro uint64) models.Passageiro {
	var passageiro models.Passageiro

	db := connectar()
	db.Where("IdPassageiro = ?", idPassageiro).Find(&passageiro)
	defer db.Close()
	fmt.Println(passageiro)

	return passageiro

}

func (PassageiroRepository) FindByAll() []models.Passageiro {
	fmt.Println("Consultando os passageiros na base")
	var passageiros []models.Passageiro
	db := connectar()
	db.Find(&passageiros)

	defer db.Close()

	return passageiros
}

func (PassageiroRepository) Update(passageiro models.Passageiro) {
	db := connectar()
	db.Save(&passageiro)
	defer db.Close()

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
