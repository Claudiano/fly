package dtos

import (
	"fly/models"
)

type PassagemDto struct {
	NumeroAcento uint64
	Voo          models.Voo
	Passageiro   models.Passageiro
}
