package routers

import (
	"fly/controllers"
	"fmt"
	"log"
	"net/http"

	_ "fly/docs"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitServer() {

	const PATH_RAIZ = "/api/v1"

	// Instancia controladores
	var vooController = controllers.VooController{}
	var passagemController = controllers.PassagemController{}
	var passageiroController = controllers.PassageiroController{}

	// Verifica serviços necessarias para funcionar

	r := mux.NewRouter()

	// endpoints de voo
	r.HandleFunc(PATH_RAIZ+"/voo", vooController.BuscarVoos).Methods("GET")
	r.HandleFunc(PATH_RAIZ+"/voo", vooController.CadastrarVoo).Methods("POST")
	r.HandleFunc(PATH_RAIZ+"/voo", vooController.ExcluirVoo).Methods("DELETE")
	r.HandleFunc(PATH_RAIZ+"/voo", vooController.AtualizarVoo).Methods("UPDATE")
	r.HandleFunc(PATH_RAIZ+"/voo/{idVoo}", vooController.BuscarVooPorId).Methods("GET")

	// endpoints de passagem
	r.HandleFunc(PATH_RAIZ+"/passagem", passagemController.BuscarPassagens).Methods("GET")
	r.HandleFunc(PATH_RAIZ+"/passagem", passagemController.CadastrarPassagem).Methods("POST")
	r.HandleFunc(PATH_RAIZ+"/passagem", passagemController.ExcluirPassagem).Methods("DELETE")
	r.HandleFunc(PATH_RAIZ+"/passagem", passagemController.AtualizarPassagem).Methods("UPDATE")
	r.HandleFunc(PATH_RAIZ+"/passagem/{idPassagem}", passagemController.BuscarPassagemPorId).Methods("GET")

	// endpoints de passageiro
	r.HandleFunc(PATH_RAIZ+"/passageiro", passageiroController.BuscarPassageiros).Methods("GET")
	r.HandleFunc(PATH_RAIZ+"/passageiro", passageiroController.CadastrarPassageiro).Methods("POST")
	r.HandleFunc(PATH_RAIZ+"/passageiro", passageiroController.AtualizarPassageiro).Methods("UPDATE")
	r.HandleFunc(PATH_RAIZ+"/passageiro/{idPassageiro}", passageiroController.BuscarPassageiroPorId).Methods("GET")

	// realizar login
	r.HandleFunc(PATH_RAIZ+"/login", passageiroController.RealizarLogin).Methods("POST")

	// adiconado documentação
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Use default options
	handler := cors.Default().Handler(r)

	// sob o servido
	fmt.Println("serviço rodando na porta 8081")
	log.Fatal(http.ListenAndServe("localhost:8081", handler))

}
