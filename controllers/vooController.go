package controllers

import (
	"encoding/json"
	"fly/dtos"
	"fly/models"
	"fly/services"
	"fly/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var vooService = services.VooService{}

type VooController struct{}

// ShowAccount godoc
// @Summary Show a account
// @Description Retorna todos os voos cadastrados
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Voo
// @Failure 404 {string} string "ok"
// @Router /voo [get]
func (VooController) BuscarVoos(w http.ResponseWriter, r *http.Request) {

	var voos []models.Voo
	voos = vooService.CarregarVoos()
	json.NewEncoder(w).Encode(voos)

}

// ShowAccount godoc
// @Summary Show a account
// @Description Retorna o voo cadastrado com base no idVoo passado
// @Accept  json
// @Produce  json
// @Param idVoo path int true "IdVoo"
// @Success 200 {object} models.Voo
// @Failure 404 {string} string "ok"
// @Router /voo/{idVoo} [get]
func (VooController) BuscarVooPorId(w http.ResponseWriter, r *http.Request) {
	var voo models.Voo
	
	// pega parametros passados
	id := chi.URLParam(r, "idVoo")

	// convertendo o idpassagem para string
	idVoo, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic("erros ao ao converter idVoo para uint64")
	}

	// serviço que retorna o voo com base no idVoo passado
	voo = vooService.CarregarVoo(idVoo)

	// adiciona como reposta o voo retornado
	json.NewEncoder(w).Encode(voo)

}

// ShowAccount godoc
// @Summary Show a account
// @Description Metodo que cadastrar voos
// @Accept  json
// @Produce  json
// @Param idVoo path int true "IdVoo"
// @Success 200 
// @Failure 404 {string} string "ok"
// @Router /voo [post]
func (VooController) CadastrarVoo(w http.ResponseWriter, r *http.Request) {
	var vooDto dtos.VooDto

	// pega o voo passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&vooDto)

	// serviço que cadastra o voo
	vooService.CadastrarVoo(vooDto)

	utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "Voo cadastrado"})
}
func (VooController) ExcluirVoo(w http.ResponseWriter, r *http.Request) {

	// voo que será adicionado
	var voo models.Voo

	// pega o voo passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&voo)

	// serviço que exclui o voo passado
	vooService.ExcluirVoo(voo)

	utils.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Voo excluido"})
}

// ShowAccount godoc
// @Summary Show a account
// @Description Metodo para atualizar um voo
// @Accept  json
// @Produce  json

// @Success 200 
// @Failure 404 {string} string "ok"
// @Router /voo [put]
func (VooController) AtualizarVoo(w http.ResponseWriter, r *http.Request) {
	
	// voo que será adicionado
	var voo models.Voo

	// pega o voo passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&voo)

	// serviço que exclui o voo passado
	vooService.AtualizarVoo(voo)

	utils.RespondwithJSON(w, http.StatusAccepted, map[string]string{"message": "Voo atualizado"})

}
