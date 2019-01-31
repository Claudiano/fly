package controllers

import (
	"encoding/json"
	"fly-go/dtos"
	"fly-go/models"
	"fly-go/services"
	"fly-go/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var passagemService = services.PassagemService{}

type PassagemController struct{}

func (PassagemController) BuscarPassagens(w http.ResponseWriter, r *http.Request) {

	var passagens []models.Passagem
	passagens, err := passagemService.CarregarPassagens()
	if err != nil {
		utils.RespondwithJSON(w, http.StatusNonAuthoritativeInfo, nil)
	} else {

		utils.RespondwithJSON(w, http.StatusOK, passagens)
	}

}

func (PassagemController) BuscarPassagemPorId(w http.ResponseWriter, r *http.Request) {
	// pega parametros passados
	id := chi.URLParam(r, "idPassagem")

	// convertendo o idpassagem para string
	idPassagem, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		panic("erros ao ao converter idPassagem para uint64")
	}
	// serviço que retorna o passagem com base no idPassagem passado
	passagem, err := passagemService.CarregarPassagem(idPassagem)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusAccepted, nil)
	} else {

		utils.RespondwithJSON(w, http.StatusOK, passagem)
	}

}
func (PassagemController) CadastrarPassagem(w http.ResponseWriter, r *http.Request) {
	var passagemDto dtos.PassagemDto

	// pega o passagem passado na requisicao
	decoder := json.NewDecoder(r.Body)
	errDec := decoder.Decode(&passagemDto)
	if errDec != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, nil)
	} else {

		// serviço que cadastra o passagem
		passagem, err := passagemService.CadastrarPassagem(passagemDto)
		if err != nil {
			utils.RespondwithJSON(w, http.StatusAccepted, nil)
		} else {
			utils.RespondwithJSON(w, http.StatusCreated, passagem)
		}

	}
}

func (PassagemController) ExcluirPassagem(w http.ResponseWriter, r *http.Request) {

	// passagem que será adicionado
	var passagem models.Passagem

	// pega o passagem passado na requisicao
	decoder := json.NewDecoder(r.Body)
	errDec := decoder.Decode(&passagem)
	if errDec != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, nil)
	} else {

		// serviço que exclui o voo passado
		res, err := passagemService.ExcluirPassagem(passagem)

		if err != nil {
			utils.RespondwithJSON(w, http.StatusAccepted, nil)
		} else {
			utils.RespondwithJSON(w, http.StatusOK, res)
		}
	}

}

func (PassagemController) AtualizarPassagem(w http.ResponseWriter, r *http.Request) {
	// passagem que será adicionado
	var passagem models.Passagem

	// pega o passagem passado na requisicao
	decoder := json.NewDecoder(r.Body)
	errDec := decoder.Decode(&passagem)
	if errDec != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, nil)
	} else {
		// serviço que exclui o passagem passado
		res, err := passagemService.AtualizarPassagem(passagem)
		if err != nil {
			utils.RespondwithJSON(w, http.StatusAccepted, nil)
		} else {
			utils.RespondwithJSON(w, http.StatusOK, res)
		}
	}

}
