package server

import (
	"log"
	"net/http"

	"github.com/fiscaluno/mu/logs"
	"github.com/fiscaluno/mu/user"
	"github.com/fiscaluno/pandorabox"

	"github.com/gorilla/mux"
)

var name string

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola, Seja bem vindo ao CRUDGo @" + name + " !!"))
}

// Listen init a http server
func Listen() {
	port := pandorabox.GetOSEnvironment("PORT", "5001")

	name = pandorabox.GetOSEnvironment("NAME", "JC")

	r := mux.NewRouter()
	r.Use(logs.LoggingMiddleware)

	user.SetRoutes(r.PathPrefix("/users").Subrouter())

	r.HandleFunc("/", handler)
	http.Handle("/", r)

	log.Println("Listen on port: " + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
