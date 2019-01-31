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

// ShowVoo godoc
// @Summary Show a voo
// @Description Retorna todos os voos cadastrados
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Voo
// @Failure 404 {string} string "ok"
// @Router /voo [get]
func (VooController) BuscarVoos(w http.ResponseWriter, r *http.Request) {

	var voos []models.Voo
	voos, err := vooService.CarregarVoos()
	if err != nil {
		utils.RespondwithJSON(w, http.StatusAccepted, nil)
	} else {
		utils.RespondwithJSON(w, http.StatusAccepted, voos)
	}

}

// ShowVoo godoc
// @Summary Show a voo
// @Description Retorna o voo cadastrado com base no idVoo passado
// @Accept  json
// @Produce  json
// @Param idVoo path int true "IdVoo"
// @Success 200 {object} models.Voo
// @Failure 404 {string} string "ok"
// @Router /voo/{idVoo} [get]
func (VooController) BuscarVooPorId(w http.ResponseWriter, r *http.Request) {
	//
	// pega parametros passados
	id := chi.URLParam(r, "idVoo")

	// convertendo o idpassagem para string
	idVoo, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic("erros ao ao converter idVoo para uint64")
	}

	// serviço que retorna o voo com base no idVoo passado
	voo, err := vooService.CarregarVoo(idVoo)

	// retorno
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, nil)
	} else {
		utils.RespondwithJSON(w, http.StatusCreated, voo)
	}

}

// ShowVoo godoc
// @Summary Show a voo
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
	bodyJ := json.NewDecoder(r.Body)
	errDec := bodyJ.Decode(&vooDto)
	if errDec != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, nil)

	} else {
		// serviço que cadastra o voo
		voo, err := vooService.CadastrarVoo(vooDto)
		if err != nil {
			utils.RespondwithJSON(w, http.StatusAccepted, nil)
		} else {
			utils.RespondwithJSON(w, http.StatusCreated, voo)

		}
	}
}

func (VooController) ExcluirVoo(w http.ResponseWriter, r *http.Request) {

	// voo que será adicionado
	var voo models.Voo

	// pega o voo passado na requisicao
	decoder := json.NewDecoder(r.Body)
	errDec := decoder.Decode(&voo)
	if errDec != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, map[string]string{"message": "Não foi possivel mapear a entidade passada no corpo da requisição"})
	} else {

		// serviço que exclui o voo passado
		err := vooService.ExcluirVoo(voo)
		if err != nil {
			utils.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Voo não excluido"})
		} else {
			utils.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Voo excluido"})
		}
	}
}

// ShowVoo godoc
// @Summary Show a voo
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
	decoder := json.NewDecoder(r.Body)
	errDec := decoder.Decode(&voo)
	if errDec != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, map[string]string{"message": "Não foi possivel mapear a entidade passada no corpo da requisição"})
	} else {

		// serviço que exclui o voo passado
		err := vooService.AtualizarVoo(voo)
		if err != nil {
			utils.RespondwithJSON(w, http.StatusAccepted, map[string]string{"message": "Voo não atualizado"})
		} else {
			utils.RespondwithJSON(w, http.StatusAccepted, map[string]string{"message": "Voo atualizado"})

		}

	}
}
