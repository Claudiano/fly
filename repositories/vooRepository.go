package repositories

import (
	"fly/dtos"
	"fly/models"
)

type VooRepository struct{}

func (VooRepository) Save(vooDto dtos.VooDto) {
	db := connectar()

	var voo models.Voo
	voo.Destino = vooDto.Destino
	voo.HoraSaida = vooDto.HoraSaida
	voo.Capacidade = vooDto.Capacidade

	db.Create(&voo)

	defer db.Close()
}

func (VooRepository) FindById(idVoo uint64) models.Voo {
	var voo models.Voo
	db := connectar()
	db.Where("idVoo = ?", idVoo).First(&voo)
	defer db.Close()

	return voo
}

func (VooRepository) FindByAll() []models.Voo {
	var voos []models.Voo

	db := connectar()
	db.Find(&voos)
	defer db.Close()

	return voos
}

func (VooRepository) Update(voo models.Voo) {
	db := connectar()
	db.Save(&voo)

	defer db.Close()
}

func (VooRepository) Delete(voo models.Voo) {
	db := connectar()
	db.Delete(&voo)
	defer db.Close()
}
