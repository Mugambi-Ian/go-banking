package middleware

import (
	"go-banking/banking-app/domain"
	"go-banking/banking-app/utils"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) AuthorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				isAuthorized := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					appError := utils.AppError{Code: http.StatusForbidden, Message: "Unauthorized"}
					utils.SendJSONResponse(w, appError.Code, appError.GetMessage())
				}
			} else {
				utils.SendJSONResponse(w, http.StatusUnauthorized, "missing token")
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}

func NewAuthMiddleware(r domain.AuthRepository) AuthMiddleware {
	return AuthMiddleware{repo: r}
}
