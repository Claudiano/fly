package controllers

import (
	"encoding/json"
	"fly/models"
	"fly/services"
	"fly/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var vooService = services.VooService{}

type VooController struct{}

func (VooController) BuscarVoos(w http.ResponseWriter, r *http.Request) {

	var voos []models.Voo
	voos = vooService.CarregarVoos()
	json.NewEncoder(w).Encode(voos)

}

func (VooController) BuscarVooPorId(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)

	// convertendo o idVoo para string
	idVoo, err := strconv.ParseUint(params["idVoo"], 10, 32)
	if err != nil {
		panic("erros ao ao converter id")
	}

	// serviço que retorna o voo com base no idVoo passado
	voo := vooService.CarregarVoo(idVoo)

	// adiciona como reposta o voo retornado
	json.NewEncoder(w).Encode(voo)

}
func (VooController) CadastrarVoo(w http.ResponseWriter, r *http.Request) {
	var voo models.Voo

	// pega o voo passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&voo)

	// serviço que cadastra o voo
	vooService.CadastrarVoo(voo)

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

func (VooController) AtualizarVoo(w http.ResponseWriter, r *http.Request) {
	// voo que será adicionado
	var voo models.Voo

	// pega o voo passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&voo)

	// serviço que exclui o voo passado
	vooService.AtualizarVoo(voo)

	utils.RespondwithJSON(w, http.StatusAccepted, map[string]string{"message": "Voo atualizado"})

}
