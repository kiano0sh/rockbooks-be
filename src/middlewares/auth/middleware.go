package auth

import (
	"context"
	"net/http"

	"gitlab.com/kian00sh/rockbooks-be/src/handlers/users"
	"gitlab.com/kian00sh/rockbooks-be/src/jwt"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Validate jwt token
			tokenStr := header
			email, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// Create user and check if user exists in db
			user := users.User{Email: email}
			id, err := users.GetUserIdByEmail(email)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id
			// Put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// And call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*users.User, error) {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	if raw == nil {
		// http.Error(w, "Invalid token", http.StatusForbidden)
		return nil, grapherrors.ReturnGQLError("دسترسی غیرمجاز!", "User has not found!")
	}
	return raw, nil
}
