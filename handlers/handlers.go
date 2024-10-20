package handlers

import (
	"log"
	"net/http"
	"os"
	"twitter/middleware"
	"twitter/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteeo mi puerto, handler y pingo a eschuchar al server*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.ChequeoBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
