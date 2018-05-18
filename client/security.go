package client

import (
	"net/http"

	"github.com/fiscaluno/pandorabox"
)

// SecurityMiddleware log all requests URI
func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		// log.Println("METHOD:", r.Method, "|", "PATH:", r.RequestURI)

		msg := pandorabox.Message{
			Content: "security",
			Status:  "OK",
			Body:    nil,
		}
		pandorabox.RespondWithJSON(w, http.StatusOK, msg)
		return
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
