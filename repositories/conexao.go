package repositories

import (
	"fly-go/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const Drive string = "postgres"
const IpServidor string = "    ec2-54-221-253-228.compute-1.amazonaws.com"
const Port int = 5432
const DbName string = "d38sqrk7l2kvjf"
const UserDatabase string = "vujqrnurjsmsdq"
const PasswordDatabase string = "    ed3a701d30a956440705558947568e4553b6f1919f33ac7ddf32a8490fa382b5"

func CriarTabelas() {
	db := connectar()

	// verifica se tabela já existem, caso não elas são criadas
	if !db.HasTable(&models.Passageiro{}) {
		//db.CreateTable(&models.Passageiro{})

		db.CreateTable(&models.Passageiro{})
		//db.Exec("CREATE TABLE passageiros(idPassageiro SERIAL PRIMARY KEY, nome VARCHAR(50), cpf VARCHAR(11), email VARCHAR(255) UNIQUE NOT NULL, senha TEXT NOT NULL);")
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
	//db, err := gorm.Open(Drive, "host=localhost port=5432 user=postgres dbname=flydb password=root@123 sslmode=disable")
	db, err := gorm.Open(Drive, "host="+IpServidor+" port=5432 user="+
		UserDatabase+" dbname="+DbName+" password="+PasswordDatabase+
		" sslmode=disable")

	if err != nil {
		panic(err)
	}

	return db
}
