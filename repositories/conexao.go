package repositories

import (
	"fly/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	Drive            string = "postgres"
	IpServidor       string = "localhost"
	Port             int    = 5432
	DbName           string = "flydb"
	UserDatabase     string = "postgres"
	PasswordDatabase string = "root@123"
)

func CriarTabelas() {
	db := connectar()

	// verifica se tabela já existem, caso não elas são criadas
	if !db.HasTable(&models.Passageiro{}) {
		db.CreateTable(&models.Passageiro{})
	}

	if !db.HasTable(&models.Voo{}) {
		db.CreateTable(&models.Voo{})
	}

	if !db.HasTable(&models.Passagem{}) {
		db.CreateTable(&models.Passagem{})
	}

	defer db.Close()
}

func connectar() *gorm.DB {

	// Abrindo conexão com o banco
	db, err := gorm.Open(Drive, "host=localhost port=5432 user=postgres dbname=flydb password=root@123 sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
