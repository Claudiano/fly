package models

type Passagem struct {
	IdPassagem   uint64 `gorm:"primary_key;column:idpassagem"`
	DataCompra   string `gorm:"column:dataCompra"`
	NumeroAcento uint64 `gorm:"column:numeroAcento"`
	IdVoo        uint64 `gorm:"column:idVoo"`
	IdPassageiro uint64 `gorm:"column:idPassageiro"`
	//	Voo   Voo    `gorm:"foreignkey:IdVoo"`
	//	Passageiro   Passageiro `gorm:"foreignkey:IdPassageiro;association_foreignkey:IdPassageiro"`
}
