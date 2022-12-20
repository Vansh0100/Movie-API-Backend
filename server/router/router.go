package router

import (
	"github.com/Vansh0100/movieapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/",controller.HomePage).Methods("GET")
	router.HandleFunc("/getAll",controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/getOne/{title}",controller.GetOneMovie).Methods("GET")
	// log.Fatal(http.ListenAndServe(":4000",router))
	return router

}
