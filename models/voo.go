package models

type Voo struct {
	IdVoo      uint64 `gorm:"PRIMARY_KEY;column:idvoo"`
	Destino    string `gorm:"column:destino"`
	HoraSaida  string `gorm:"column:horaSaida"`
	Capacidade uint64 `gorm:"column:capacidade"`
}
