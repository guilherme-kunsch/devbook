package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func ToGenerate() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configure(r)

}
