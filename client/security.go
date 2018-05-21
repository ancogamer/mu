package client

import (
	"log"
	"net/http"

	"github.com/fiscaluno/pandorabox"
)

// SecurityMiddleware log all requests URI
func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		msg := pandorabox.Message{
			Content: "Security",
			Status:  "Unauthorized",
			Body:    nil,
		}

		clientID := pandorabox.GetOSEnvironment("CLIENT_ID", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImNsaWVudCI6IkpDIn0sImlzcyI6Im11In0.spG-5_yioatRlT_em0i0p7XUP3J5Ne9XxiZL-GRK16Y")

		if r.Header.Get("X-Client-ID") == "" {
			log.Println("X-Client-ID is NULL")
			pandorabox.RespondWithJSON(w, http.StatusUnauthorized, msg)
			return
		}

		if r.Header.Get("X-Client-ID") != clientID {
			log.Println("X-Client-ID " + r.Header.Get("X-Client-ID") + " don't matching")
			pandorabox.RespondWithJSON(w, http.StatusUnauthorized, msg)
			return
		}

		log.Println("X-Client-ID is " + r.Header.Get("X-Client-ID"))

		next.ServeHTTP(w, r)
	})
}
