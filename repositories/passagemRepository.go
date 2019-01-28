package repositories

import (
	"fly/models"
)

type PassagemRepository struct{}

func (PassagemRepository) Save(passagem models.Passagem) {
	db := connectar()
	db.Create(passagem)
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
