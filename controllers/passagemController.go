package controllers

import (
	"encoding/json"
	"fly/dtos"
	"fly/models"
	"fly/services"
	"fly/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var passagemService = services.PassagemService{}

type PassagemController struct{}

func (PassagemController) BuscarPassagens(w http.ResponseWriter, r *http.Request) {

	var passagens []models.Passagem
	passagens = passagemService.CarregarPassagens()
	json.NewEncoder(w).Encode(passagens)

}

func (PassagemController) BuscarPassagemPorId(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)

	// convertendo o idpassagem para string
	idPassagem, err := strconv.ParseUint(params["idPassagem"], 10, 64)
	if err != nil {
		panic("erros ao ao converter idPassagem para uint64")
	}

	// serviço que retorna o passagem com base no idPassagem passado
	passagem := passagemService.CarregarPassagem(idPassagem)

	// adiciona como reposta o voo retornado
	json.NewEncoder(w).Encode(passagem)

}
func (PassagemController) CadastrarPassagem(w http.ResponseWriter, r *http.Request) {
	var passagemDto dtos.PassagemDto

	// pega o passagem passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&passagemDto)

	// serviço que cadastra o passagem
	passagemService.CadastrarPassagem(passagemDto)

	utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "Passagem cadastrado"})
}
func (PassagemController) ExcluirPassagem(w http.ResponseWriter, r *http.Request) {

	// passagem que será adicionado
	var passagem models.Passagem

	// pega o passagem passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&passagem)

	// serviço que exclui o voo passado
	passagemService.ExcluirPassagem(passagem)

	utils.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Passagem excluida"})
}

func (PassagemController) AtualizarPassagem(w http.ResponseWriter, r *http.Request) {
	// passagem que será adicionado
	var passagem models.Passagem

	// pega o passagem passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&passagem)

	// serviço que exclui o passagem passado
	passagemService.AtualizarPassagem(passagem)

	utils.RespondwithJSON(w, http.StatusAccepted, map[string]string{"message": "Passagem atualizado"})

}
