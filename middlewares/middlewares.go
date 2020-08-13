package middlewares

import (
	"net/http"

	"github.com/urfave/negroni"
)

// ApplicationJSON - Defini ContentType como Json
func ApplicationJSON() negroni.Handler {

	// Retorna Comportamento
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		// Seta Content-Type como Json
		rw.Header().Set("Content-Type", "application/json")
		next(rw, r)
	})
}

// EnableCors - Habilita Acesso ao Sistema
func EnableCors() negroni.Handler {

	// Retorna Comportamento
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		// Seta Acesso ao Sistema
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		next(rw, r)
	})
}
