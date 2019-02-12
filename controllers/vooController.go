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

var vooService = services.VooService{}

type VooController struct{}

type Ivoo interface {
	BuscarVoos()
}

// @Title BuscarVoos
// @Summary Show a voo
// @Description Retorna todos os voos cadastrados
// @Tags voo
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Voo
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /voo [get]
func (VooController) BuscarVoos(w http.ResponseWriter, r *http.Request) {

	var voos []models.Voo
	voos, err := vooService.CarregarVoos()
	if err != nil {
		utils.RespondwithJSON(w, http.StatusNotFound, nil)
	} else {
		utils.RespondwithJSON(w, http.StatusAccepted, voos)
	}

}

// ShowVoo godoc
// @Summary Show a voo
// @Description Retorna o voo cadastrado com base no idVoo passado
// @Tags voo
// @Accept  json
// @Produce  json
// @Param idVoo path int true "IdVoo"
// @Success 200 {object} models.Voo
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /voo/{idVoo} [get]
func (VooController) BuscarVooPorId(w http.ResponseWriter, r *http.Request) {
	//
	// pega parametros passados
	id := chi.URLParam(r, "idVoo")

	// convertendo o idpassagem para string
	idVoo, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, nil)
		panic("erros ao ao converter idVoo para uint64")
	} else {

		// serviço que retorna o voo com base no idVoo passado
		voo, err := vooService.CarregarVoo(idVoo)

		// retorno
		if err != nil {
			utils.RespondwithJSON(w, http.StatusNotFound, nil)
		} else {
			utils.RespondwithJSON(w, http.StatusCreated, voo)
		}
	}

}

// ShowVoo godoc
// @Summary Show a voo
// @Description Metodo que cadastrar voos
// @Tags voo
// @Accept  json
// @Produce  json
// @Param idVoo path int true "IdVoo"
// @Success 200 {object} models.Voo
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
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
			utils.RespondwithJSON(w, http.StatusNotFound, nil)
		} else {
			utils.RespondwithJSON(w, http.StatusCreated, voo)

		}
	}
}

// ShowVoo godoc
// @Summary Show a voo
// @Description Metodo que cadastrar voos
// @Tags voo
// @Accept  json
// @Produce  json
// @Param idVoo path int true "IdVoo"
// @Param destino path string true "IdVoo"
// @Param horaSaida path string true "IdVoo"
// @Param capacidade path int true "IdVoo"
// @Success 200 {object} models.Voo
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /voo [delete]
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
			utils.RespondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Voo não excluido"})
		} else {
			utils.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Voo excluido"})
		}
	}
}

// ShowVoo godoc
// @Summary Show a voo
// @Description Metodo para atualizar um voo
// @Tags voo
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Resource /voo
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
			utils.RespondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Voo não atualizado"})
		} else {
			utils.RespondwithJSON(w, http.StatusAccepted, map[string]string{"message": "Voo atualizado"})

		}

	}
}
