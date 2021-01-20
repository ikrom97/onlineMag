package middlewares

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"onlineMag/token"
)

func Authorized() func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
			bearerToken := r.Header.Get("Authorization")
			Token := bearerToken[len("Bearer "):]
			claims := token.ParseToken(Token)
			if claims.Login == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			next(w, r, pr)
		}
	}
}
