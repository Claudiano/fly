package repositories

import (
	"fly/dtos"
	"fly/models"
)

type PassagemRepository struct{}

func (PassagemRepository) Save(passagemDto dtos.PassagemDto) {
	var passagem models.Passagem

	db := connectar()

	passagem.Passageiro = passagemDto.Passageiro
	passagem.DataCompra = "12/12/2019" //formatar data
	passagem.IdVoo = passagemDto.Voo.IdVoo
	passagem.Voo = passagemDto.Voo
	passagem.IdPassageiro = passagemDto.Passageiro.IdPassageiro
	passagem.Passageiro = passagemDto.Passageiro

	//db.Create(passagem)
	db.Save(passagem)
	defer db.Close()

}

func (PassagemRepository) FindById(idPassagem uint64) models.Passagem {

	var passagem models.Passagem
	db := connectar()
	db.Where("idPassagem = ?", idPassagem).First(&passagem)
	defer db.Close()

	return passagem
}

func (PassagemRepository) FindByAll() []models.Passagem {

	var passagens []models.Passagem

	db := connectar()
	db.Find(&passagens)
	defer db.Close()

	return passagens
}

func (PassagemRepository) Update(passagem models.Passagem) {
	db := connectar()
	db.Save(passagem)
	defer db.Close()
}

func (PassagemRepository) Delete(passagem models.Passagem) {
	db := connectar()
	db.Delete(passagem)
	defer db.Close()

}
