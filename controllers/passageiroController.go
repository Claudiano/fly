package controllers

import (
	"encoding/json"
	"fly/dtos"
	"fly/models"
	"fly/services"
	"fly/settings"
	"fly/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var passageiroService = services.PassageiroService{}

type PassageiroController struct{}

func (PassageiroController) BuscarPassageiros(w http.ResponseWriter, r *http.Request) {

	var passageiros []models.Passageiro
	passageiros = passageiroService.BuscarPassageiros()
	json.NewEncoder(w).Encode(passageiros)

}

func (PassageiroController) BuscarPassageiroPorId(w http.ResponseWriter, r *http.Request) {

	// pega parametros passados
	id := chi.URLParam(r, "idPassageiro")

	// convertendo o idpassagem para string
	idPassageiro, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic("erros ao ao converter idPassageiro para uint64")
	}

	// serviço que retorna o passagem com base no idPassagem passado
	passageiro := passageiroService.BuscarPassageiro(idPassageiro)

	// adiciona como reposta o voo retornado
	json.NewEncoder(w).Encode(passageiro)

}
func (PassageiroController) CadastrarPassageiro(w http.ResponseWriter, r *http.Request) {
	var passageiro dtos.PassageiroDto

	// pega o passageiro passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&passageiro)

	// serviço que cadastra o passageiro
	passageiroService.CadastrarPassageiro(passageiro)

	utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "Passageiro cadastrado"})
}

func (PassageiroController) AtualizarPassageiro(w http.ResponseWriter, r *http.Request) {

	var passageiro models.Passageiro

	// pega o passageiro passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&passageiro)

	// serviço que cadastra o passageiro
	passageiroService.AtualizarPassageiro(passageiro)

	utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "Passageiro Atualizado"})

}

func (PassageiroController) RealizarLogin(w http.ResponseWriter, r *http.Request) {
	var credenciaisDto dtos.CredenciaisDto
	// pega o passageiro passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&credenciaisDto)

	// serviço que cadastra o passageiro
	passageiro := passageiroService.BuscarPassageiroLogin(credenciaisDto)

	// adiciona como reposta o voo retornado
	//	json.NewEncoder(w).Encode(passageiro)
	//	utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "Passageiro cadastrado"})

	fmt.Println(credenciaisDto)

	if credenciaisDto.Email == passageiro.Email && credenciaisDto.Senha == passageiro.Senha {

		token, err := settings.GenerateJWT(passageiro)
		if err != nil {
			log.Fatal(err)
		}

		result := models.ResponseToken{token}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)

	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("usuario invalido")
	}
}
