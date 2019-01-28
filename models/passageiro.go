package models

type Passageiro struct {
	IdPassageiro uint64 `gorm:PRIMARY_KEY;AUTO_INCREMENT`
	Nome         string `gorm:"column:nome"`
	Cpf          string `gorm:"column:cpf"`
	Email        string `gorm:"column:email"`
	Senha        string `gorm:"column:senha"`
}
