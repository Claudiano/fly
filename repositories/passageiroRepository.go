package repositories

import (
	"fly/dtos"
	"fly/models"
	"fmt"
)

type PassageiroRepository struct {
}

func (PassageiroRepository) Save(passageiro models.Passageiro) {
	db := connectar()
	db.Create(passageiro)
	defer db.Close()
}

func (PassageiroRepository) FindById(idPassageiro uint64) models.Passageiro {
	var passageiro models.Passageiro

	db := connectar()
	db.Where("idPassageiro = ?", idPassageiro).Find(&passageiro)
	defer db.Close()

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
	db.Save(passageiro)
	defer db.Close()

}

func (PassageiroRepository) Delete(passageiro models.Passageiro) {
	db := connectar()

	db.Delete(passageiro)

	defer db.Close()

}

func (PassageiroRepository) FIndByPassageiro(passageiroDto dtos.PassageiroDto) models.Passageiro {
	var passageiro models.Passageiro
	db := connectar()

	db.Where("email =  ? and senha = ?", passageiroDto.Email, passageiroDto.Senha).First(&passageiro)

	defer db.Close()

	return passageiro

}
