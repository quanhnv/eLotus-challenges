package middlewares

import (
	"net/http"
	"github.com/quanhnv/eLotus-challenges/auth"
)

func CheckJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := jwtHelper.Verify(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
