package repositories

import (
	"fly-go/dtos"
	"fly-go/models"
)

type VooRepository struct{}

func (VooRepository) Save(vooDto dtos.VooDto) (models.Voo, error) {
	db := connectar()

	var voo models.Voo
	voo.Destino = vooDto.Destino
	voo.HoraSaida = vooDto.HoraSaida
	voo.Capacidade = vooDto.Capacidade

	err := db.Create(&voo).First(&voo).Error

	defer db.Close()

	if err != nil {
		return voo, err
	}
	return voo, nil
}

func (VooRepository) FindById(idVoo uint64) (models.Voo, error) {
	var voo models.Voo
	db := connectar()
	err := db.Where("idVoo = ?", idVoo).First(&voo).Error
	defer db.Close()

	if err != nil {
		return voo, err
	}
	return voo, nil
}

func (VooRepository) FindByAll() ([]models.Voo, error) {
	var voos []models.Voo

	db := connectar()
	err := db.Find(&voos).Error
	defer db.Close()

	if err != nil {
		return voos, err
	}
	return voos, nil
}

func (VooRepository) Update(voo models.Voo) error {
	db := connectar()
	err := db.Save(&voo).Error

	defer db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (VooRepository) Delete(voo models.Voo) error {
	db := connectar()
	err := db.Delete(&voo).Error
	defer db.Close()

	if err != nil {
		return err
	}
	return err
}
