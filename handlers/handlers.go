package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/luisesteban91/curso_go_api_twiter/middleware"
	"github.com/luisesteban91/curso_go_api_twiter/routers"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/miperfil", middleware.ChequeoBD(middleware.ValidateJWT(routers.MiPerfil))).Methods("POST")

	PORT := os.Getenv("PORT") //obetener variable de entorno

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) //establecer todos los ṕermisos de cors
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
