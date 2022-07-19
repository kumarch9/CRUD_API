package routing

import (
	"log"
	"net/http"
	hd "webapiingo/handlers"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	route := mux.NewRouter()
	route.HandleFunc("/worker", hd.CreateWorker).Methods("POST")
	route.HandleFunc("/worker", hd.GetWorker).Methods("GET")
	route.HandleFunc("/worker/{id}", hd.GetWorkerByID).Methods("GET")
	route.HandleFunc("/worker/{id}", hd.UpdateWorker).Methods("PUT")
	route.HandleFunc("/worker/{id}", hd.DeleteWorker).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8055", route))

}
