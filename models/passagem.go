package models

type Passagem struct {
	IdPassagem   uint64 `gorm:"primary_key;column:idPassagem"`
	DataCompra   string `gorm:"column:dataCompra"`
	NumeroAcento uint64 `gorm:"column:numeroAcento"`

	IdVoo uint64 `gorm:"column:idVoo"`
	Voo   Voo    `gorm:"foreignkey:IdVoo;association_foreignkey:IdVoo"`

	IdPassageiro uint64     `gorm:"column:idPassageiro"`
	Passageiro   Passageiro `gorm:"foreignkey:IdPassageiro;association_foreignkey:IdPassageiro"`
}
