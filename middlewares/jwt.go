package middlewares

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"onlineMag/token"
	"strings"
	"time"
)

func JWT() func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
			bearerToken := r.Header.Get("Authorization")
			if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			Token := bearerToken[len("Bearer "):]
			claims := token.ParseToken(Token)
			if claims.ExpiresAt < time.Now().Unix() {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			next(w, r, pr)
		}
	}
}
