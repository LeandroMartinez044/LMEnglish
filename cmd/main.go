package main

import (
	"lmenglish/cmd/dependencies"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	d := dependencies.Init()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/record", d.RecordHandler.Record).Methods("POST")

	handler := cors.Default().Handler(myRouter)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
