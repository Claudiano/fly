package routers

import (
	"fly-go/controllers"
	"fly-go/settings"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	_ "fly-go/docs"

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

	routesAuth := r.Group(nil)
	routesAuth.Use(settings.AuthMiddleware)

	// realizar login
	r.Post(PATH_RAIZ+"/login", passageiroController.RealizarLogin)

	// endpoints de voo
	routesAuth.Get(PATH_RAIZ+"/voo", vooController.BuscarVoos)
	routesAuth.Post(PATH_RAIZ+"/voo", vooController.CadastrarVoo)
	routesAuth.Delete(PATH_RAIZ+"/voo", vooController.ExcluirVoo)
	routesAuth.Put(PATH_RAIZ+"/voo", vooController.AtualizarVoo)
	routesAuth.Get(PATH_RAIZ+"/voo/{idVoo}", vooController.BuscarVooPorId)

	// endpoints de passagem
	routesAuth.Get(PATH_RAIZ+"/passagem", passagemController.BuscarPassagens)
	routesAuth.Post(PATH_RAIZ+"/passagem", passagemController.CadastrarPassagem)
	routesAuth.Delete(PATH_RAIZ+"/passagem", passagemController.ExcluirPassagem)
	routesAuth.Put(PATH_RAIZ+"/passagem", passagemController.AtualizarPassagem)
	routesAuth.Get(PATH_RAIZ+"/passagem/{idPassagem}", passagemController.BuscarPassagemPorId)

	// endpoints de passageiro
	r.Post(PATH_RAIZ+"/passageiro", passageiroController.CadastrarPassageiro)
	routesAuth.Get(PATH_RAIZ+"/passageiro", passageiroController.BuscarPassageiros)
	routesAuth.Put(PATH_RAIZ+"/passageiro", passageiroController.AtualizarPassageiro)
	routesAuth.Get(PATH_RAIZ+"/passageiro/{idPassageiro}", passageiroController.BuscarPassageiroPorId)

	// adiconado documentação
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Use default options
	handler := cors.Default().Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// sob o servido
	fmt.Println("serviço rodando na porta", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
