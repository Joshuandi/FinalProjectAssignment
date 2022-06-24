package middleware

import (
	"FinalProjectAssignment/auth"
	"FinalProjectAssignment/model"
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var ctxKey = &contextKey{"users"}

type contextKey struct {
	data string
}

func AuthIsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.Contains(header, "Bearer") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := ""
		lenToken := strings.Split(header, " ")
		if len(lenToken) == 2 {
			tokenString = lenToken[1]
		}

		tokenc, err := auth.ValidationToken(tokenString)
		if err != nil {
			http.Error(w, "Token UnAuthorized", http.StatusNotAcceptable)
			return
		}

		claims, ok := tokenc.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Token UnAuthorized", http.StatusNotAcceptable)
			return
		}

		user_id := claims["id"].(string)
		id, _ := strconv.Atoi(user_id)
		users := model.User{User_id: id}
		ctx := context.WithValue(r.Context(), ctxKey, &users)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

//find user from context, need middleware to run
func runContext(ctx context.Context) model.User {
	rawData, _ := ctx.Value(ctxKey).(model.User)
	return rawData
}
