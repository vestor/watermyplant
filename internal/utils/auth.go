package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/vestor/watermyplant/internal/constants"
	"github.com/vestor/watermyplant/internal/log"
	"net/http"

	p "github.com/vestor/watermyplant/internal/pogos"
)

var l = log.Get()
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var cookie, _ = r.Cookie(constants.SESSION_COOKIE_NAME)

		if cookie == nil {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(p.Exception{Message: "Missing Authorization"})
			return
		}
		tk := &p.Token{}

		_, err := jwt.ParseWithClaims(cookie.Value, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.SECRET), nil
		})

		fmt.Printf("Token %+v \n", tk.Username)

		if err != nil {
			l.Println(err)
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(p.Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
