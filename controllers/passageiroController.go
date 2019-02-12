package controllers

import (
	"encoding/json"
	"fly-go/dtos"
	"fly-go/models"
	"fly-go/services"
	"fly-go/settings"
	"fly-go/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var passageiroService = services.PassageiroService{}

type PassageiroController struct{}

// ShowPassageiro godoc
// @Summary Show a passageiro
// @Description Retorna todos os passageiros cadastrados
// @Tags Passageiro
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Passageiro
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /passageiro [get]
func (PassageiroController) BuscarPassageiros(w http.ResponseWriter, r *http.Request) {

	var passageiros []models.Passageiro
	passageiros, err := passageiroService.BuscarPassageiros()
	if err != nil {
		fmt.Println(err)
		utils.RespondwithJSON(w, http.StatusNotFound, nil)
	} else {
		utils.RespondwithJSON(w, http.StatusBadRequest, passageiros)
	}

}

// ShowPassageiro godoc
// @Summary Show a passageiro
// @Description Retorna o passageiros cadastrado com base no idPassageiro passado
// @Tags Passageiro
// @Accept  json
// @Produce  json
// @Param idVoo path int true "IdPassageiro"
// @Success 200 {object} models.Passageiro
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /passageiro/{idPassageiro} [get]
func (PassageiroController) BuscarPassageiroPorId(w http.ResponseWriter, r *http.Request) {

	// pega parametros passados
	id := chi.URLParam(r, "idPassageiro")

	// convertendo o idpassagem para string
	idPassageiro, err := strconv.ParseUint(id, 10, 32)
	if err != nil {

		utils.RespondwithJSON(w, http.StatusBadRequest, nil)
		panic("erros ao ao converter idPassageiro para uint64")
	}

	// serviço que retorna o passagem com base no idPassagem passado
	passageiro, err := passageiroService.BuscarPassageiro(idPassageiro)
	if err != nil {
		fmt.Println(err)
		utils.RespondwithJSON(w, http.StatusNotFound, nil)
	} else {
		utils.RespondwithJSON(w, http.StatusOK, passageiro)
	}

}

// ShowPassageiro godoc
// @Summary Show a passageiro
// @Description Metodo que cadastrar passageiro
// @Tags Passageiro
// @Accept  json
// @Produce  json
// @Param Passageiro body dtos.PassageiroDto true "Passageiro"
// @Success 200 {string} string "Passageiro cadastrado"
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /passageiro [post]
func (PassageiroController) CadastrarPassageiro(w http.ResponseWriter, r *http.Request) {
	var passageiro dtos.PassageiroDto

	// pega o passageiro passado na requisicao
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&passageiro)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, map[string]string{"message": "Passageiro não cadastrado"})
	} else {
		// serviço que cadastra o passageiro
		_, err := passageiroService.CadastrarPassageiro(passageiro)
		if err != nil {
			utils.RespondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Passageiro nao cadastrado"})
		} else {
			utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "Passageiro cadastrado"})
		}
	}

}

// ShowVoo godoc
// @Summary Show a voo
// @Description Metodo para atualizar um voo
// @Tags Passageiro
// @Accept  json
// @Produce  json
// @Param Passageiro body models.Passageiro true "Passageiro"
// @Success 200 {string} string "Passageiro atualizado"
// @Failure 400 {string} string "Requisição invalida"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /passageiro [put]
func (PassageiroController) AtualizarPassageiro(w http.ResponseWriter, r *http.Request) {

	var passageiro models.Passageiro

	// pega o passageiro passado na requisicao
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&passageiro)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest, map[string]string{"message": "Passageiro não atualizado"})
	} else {
		// serviço que cadastra o passageiro
		err := passageiroService.AtualizarPassageiro(passageiro)
		if err != nil {
			utils.RespondwithJSON(w, http.StatusNotFound, map[string]string{"message": "Passageiro não atualizado"})
		} else {
			utils.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Passageiro Atualizado"})
		}
	}

}

// Show Passageiro godoc
// @Summary realizar login
// @Description verifica na base se existe as credenciais passadas, e retorna um token caso exista.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param Passageiro body dtos.CredenciaisDto true "Login"
// @Success 200 {object} models.ResponseToken true
// @Failure 400 {string} string "Requisição invalida"
// @Failure 404 {string} string "Nenhum registro encontrado."
// @Router /login [post]
func (PassageiroController) RealizarLogin(w http.ResponseWriter, r *http.Request) {
	var credenciaisDto dtos.CredenciaisDto
	// pega o passageiro passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&credenciaisDto)

	// serviço que cadastra o passageiro
	passageiro := passageiroService.BuscarPassageiroLogin(credenciaisDto)

	if credenciaisDto.Email == passageiro.Email && credenciaisDto.Senha == passageiro.Senha {

		token, err := settings.GenerateJWT(passageiro)
		if err != nil {
			log.Fatal(err)
		}

		result := models.ResponseToken{token}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)

	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("usuario invalido")
	}
}
