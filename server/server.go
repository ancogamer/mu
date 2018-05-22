package server

import (
	"log"
	"net/http"

	"github.com/fiscaluno/mu/client"
	"github.com/fiscaluno/mu/logs"
	"github.com/fiscaluno/mu/user"
	"github.com/fiscaluno/pandorabox"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var name string

func handlerHi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola, Seja bem vindo ao CRUDGo @" + name + " !!"))
}

// Listen init a http server
func Listen() {
	port := pandorabox.GetOSEnvironment("PORT", "5001")

	name = pandorabox.GetOSEnvironment("NAME", "JC")

	r := mux.NewRouter()
	r.Use(logs.LoggingMiddleware)
	r.Use(client.SecurityMiddleware)

	user.SetRoutes(r.PathPrefix("/users").Subrouter())

	r.HandleFunc("/", handlerHi)
	http.Handle("/", r)

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Client-ID", "Content-Type"})
	methodsOk := handlers.AllowedMethods([]string{"*"})

	log.Println("Listen on port: " + port)
	if err := http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
		log.Fatal(err)
	}
}
