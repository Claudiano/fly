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

	"github.com/gorilla/mux"
)

var passageiroService = services.PassageiroService{}

type PassageiroController struct{}

func (PassageiroController) BuscarPassageiros(w http.ResponseWriter, r *http.Request) {

	var passageiros []models.Passageiro
	passageiros = passageiroService.BuscarPassageiros()
	json.NewEncoder(w).Encode(passageiros)

}

func (PassageiroController) BuscarPassageiroPorId(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)

	// convertendo o idpassagem para string
	idPassageiro, err := strconv.ParseUint(params["idPassageiro"], 10, 64)
	if err != nil {
		panic("erros ao ao converter idPassageiro para uint64")
	}

	// serviço que retorna o passagem com base no idPassagem passado
	passageiro := passageiroService.BuscarPassageiro(idPassageiro)

	// adiciona como reposta o voo retornado
	json.NewEncoder(w).Encode(passageiro)

}
func (PassageiroController) CadastrarPassageiro(w http.ResponseWriter, r *http.Request) {
	var passageiro models.Passageiro

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
	var passageiroDto dtos.PassageiroDto
	var passageiro = models.Passageiro{Cpf: "123456789", Email: "teste", IdPassageiro: 1, Senha: "123456"}
	// pega o passageiro passado na requisicao
	_ = json.NewDecoder(r.Body).Decode(&passageiroDto)

	// serviço que cadastra o passageiro
	//	passageiro := passageiroService.BuscarPassageiroLogin(passageiroDto)

	// adiciona como reposta o voo retornado
	//	json.NewEncoder(w).Encode(passageiro)
	//	utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "Passageiro cadastrado"})

	// Demo - in real case scenario you'd check this against your database
	fmt.Println(passageiroDto.Email + " : " + passageiroDto.Senha)

	if passageiroDto.Email == "claudiano" && passageiroDto.Senha == "claudiano" {
		fmt.Println("usuario logado")

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
