package routers

import (
	"fly/controllers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	_ "fly/docs"

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

	r := chi.NewRouter()

	// endpoints de voo
	r.Get(PATH_RAIZ+"/voo", vooController.BuscarVoos)
	r.Post(PATH_RAIZ+"/voo", vooController.CadastrarVoo)
	r.Delete(PATH_RAIZ+"/voo", vooController.ExcluirVoo)
	r.Put(PATH_RAIZ+"/voo", vooController.AtualizarVoo)
	r.Get(PATH_RAIZ+"/voo/{idVoo}", vooController.BuscarVooPorId)

	// endpoints de passagem
	r.Get(PATH_RAIZ+"/passagem", passagemController.BuscarPassagens)
	r.Post(PATH_RAIZ+"/passagem", passagemController.CadastrarPassagem)
	r.Delete(PATH_RAIZ+"/passagem", passagemController.ExcluirPassagem)
	r.Put(PATH_RAIZ+"/passagem", passagemController.AtualizarPassagem)
	r.Get(PATH_RAIZ+"/passagem/{idPassagem}", passagemController.BuscarPassagemPorId)

	// endpoints de passageiro
	r.Get(PATH_RAIZ+"/passageiro", passageiroController.BuscarPassageiros)
	r.Post(PATH_RAIZ+"/passageiro", passageiroController.CadastrarPassageiro)
	r.Put(PATH_RAIZ+"/passageiro", passageiroController.AtualizarPassageiro)
	r.Get(PATH_RAIZ+"/passageiro/{idPassageiro}", passageiroController.BuscarPassageiroPorId)

	// realizar login
	r.Post(PATH_RAIZ+"/login", passageiroController.RealizarLogin)

	// adiconado documentação
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Use default options
	handler := cors.Default().Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	// sob o servido
	fmt.Println("serviço rodando na porta", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
