package models

import (
	"time"
)

type Voo struct {
	IdVoo      uint64    `gorm:PRIMARY_KEY;AUTO_INCREMENT`
	Destino    string    `gorm:"column:destino"`
	HoraSaida  time.Time `gorm:"column:horaSaida"`
	Capacidade uint64    `gorm:"column:capacidade"`
}
