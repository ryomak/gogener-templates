package middleware

import (
	"net/http"
	"strings"

	"[[.ModName]]/infrastructure/firebase/auth"
	"[[.ModName]]/pkg/context"
	"[[.ModName]]/pkg/log"
	"[[.ModName]]/presentation/api/v1/resource/factory"
)

type AuthMiddleware struct {
	authClient auth.Client
}

func NewAuth(authClient auth.Client) *AuthMiddleware {
	return &AuthMiddleware{
		authClient,
	}
}

func (a AuthMiddleware) Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		ctx := r.Context()
		uid, err := a.authClient.VerifyIDToken(ctx, idToken)
		if err != nil {
			log.Errorf(ctx, "Auth.VerifyIDToken %v", err)
			factory.ErrorJSON(w, err, http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithAuthID(ctx, *uid))

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
