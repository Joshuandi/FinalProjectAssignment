package middleware

import (
	"FinalProjectAssignment/auth"
	"FinalProjectAssignment/config"
	"FinalProjectAssignment/model"
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var Key = &contextData{"users"}
var cfg *config.Config

type contextData struct {
	data string
}

//find user from context, need middleware to run
func RunUser(ctx context.Context) *model.User {
	rawData, _ := ctx.Value(Key).(*model.User)
	return rawData
}

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authoriHeader := r.Header.Get("Authorization")
		if !strings.Contains(authoriHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}
		tokenString := strings.Replace(authoriHeader, "Bearer ", "", -1)
		//fmt.Println("ini token string:", tokenString)
		tokenValid, err := auth.ValidationToken(tokenString)
		//fmt.Println("ini token valid:", tokenValid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		claims, ok := tokenValid.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Token UnAuthorized", http.StatusNotAcceptable)
			return
		}
		//fmt.Println("ini claims:", claims)
		user_id := claims["User_id"].(string)
		//fmt.Println("ini user_id:", user_id)
		user_id_, _ := strconv.Atoi(user_id)
		userss := model.User{User_id: user_id_}
		//fmt.Println("ini userss middleware:", userss)
		ctx := context.WithValue(r.Context(), Key, &userss)
		//fmt.Println("ini ctx middle:", ctx)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
