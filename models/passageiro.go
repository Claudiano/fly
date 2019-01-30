package models

type Passageiro struct {
	IdPassageiro uint64 `gorm:"primary_key;column:idpassageiro"`
	Nome         string `gorm:"column:nome"`
	Cpf          string `gorm:"column:cpf"`
	Email        string `gorm:"column:email"`
	Senha        string `gorm:"column:senha"`
}
