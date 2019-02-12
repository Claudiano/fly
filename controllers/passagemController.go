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

// ShowVoo godoc
// @Summary Show a passagens
// @Description Metodo para buscar todas as passagens
// @Tags Passagem
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Passagem
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /passagem [get]
func (PassagemController) BuscarPassagens(w http.ResponseWriter, r *http.Request) {

	var passagens []models.Passagem
	passagens, err := passagemService.CarregarPassagens()
	if err != nil {
		utils.RespondwithJSON(w, http.StatusNotFound, []models.Passagem{})
	} else {
		if len(passagens) == 0 {
			utils.RespondwithJSON(w, http.StatusNotFound, []models.Passagem{})
		} else {
			utils.RespondwithJSON(w, http.StatusOK, passagens)
		}
	}

}

// ShowVoo godoc
// @Summary Show a passagens
// @Description Metodo para buscar passagem passando o id da passagem como referencia
// @Tags Passagem
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Passagem
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /passagem [get]
func (PassagemController) BuscarPassagemPorId(w http.ResponseWriter, r *http.Request) {
	// pega parametros passados
	id := chi.URLParam(r, "idPassagem")

	// convertendo o idpassagem para string
	idPassagem, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, nil)
	}
	// serviço que retorna o passagem com base no idPassagem passado
	passagem, err := passagemService.CarregarPassagem(idPassagem)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusAccepted, nil)
	} else {

		utils.RespondwithJSON(w, http.StatusOK, passagem)
	}

}

// ShowVoo godoc
// @Summary Show a passagens
// @Description Metodo para registrar uma pasagem
// @Tags Passagem
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Passagem
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /passagem [post]
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
			utils.RespondwithJSON(w, http.StatusNotImplemented, models.Passagem{})
		} else {
			utils.RespondwithJSON(w, http.StatusCreated, passagem)
		}

	}
}

// ShowVoo godoc
// @Summary Show a passagens
// @Description Metodo para excluir uma pasagem
// @Tags Passagem
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Passagem
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /passagem [delete]
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
			utils.RespondwithJSON(w, http.StatusNotFound, map[string]string{"messaege": "Não foi possivel salvar excluir a passagem."})
		} else {
			utils.RespondwithJSON(w, http.StatusOK, res)
		}
	}

}

// ShowVoo godoc
// @Summary Show a passagens
// @Description Metodo atualiza uma passagem
// @Tags Passagem
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Passagem
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /passagem [put]
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
			utils.RespondwithJSON(w, http.StatusNotFound, passagem)
		} else {
			utils.RespondwithJSON(w, http.StatusOK, res)
		}
	}

}
