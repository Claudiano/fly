package repositories

import (
	"fly-go/dtos"
	"fly-go/models"
)

type PassagemRepository struct{}

func (PassagemRepository) Save(passagemDto dtos.PassagemDto) (models.Passagem, error) {
	var passagem models.Passagem

	db := connectar()

	//passagem.Passageiro = passagemDto.Passageiro
	//passagem.Voo = passagemDto.Voo
	//passagem.Passageiro = passagemDto.Passageiro
	passagem.DataCompra = "12/12/2019" //formatar data
	passagem.IdVoo = passagemDto.IdVoo
	passagem.IdPassageiro = passagemDto.IdPassageiro
	passagem.NumeroAcento = passagemDto.NumeroAcento

	//db.Create(passagem)
	err := db.Save(&passagem).First(&passagem).Error
	defer db.Close()

	if err != nil {
		return passagem, err
	}
	return passagem, nil
}

func (PassagemRepository) FindById(idPassagem uint64) (models.Passagem, error) {

	var passagem models.Passagem

	db := connectar()
	err := db.Where("idpassagem = ?", idPassagem).First(&passagem).Error
	defer db.Close()

	return passagem, err
}

func (PassagemRepository) FindByAll() ([]models.Passagem, error) {

	var passagens []models.Passagem

	db := connectar()
	err := db.Find(&passagens).Error
	defer db.Close()

	return passagens, err
}

func (PassagemRepository) Update(passagem models.Passagem) (models.Passagem, error) {
	var passagemUp models.Passagem

	db := connectar()
	err := db.Save(&passagem).Error
	defer db.Close()

	return passagemUp, err
}

func (PassagemRepository) Delete(passagem models.Passagem) (models.Passagem, error) {
	var passagemDeletada models.Passagem

	db := connectar()
	err := db.Delete(passagem).First(&passagemDeletada).Error
	defer db.Close()

	return passagemDeletada, err

}
