package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

// User is the user model
type User struct {
	UUID  string
	Email string
	Link  struct {
		Roles     []string
		Grantings []string
	}
}

func ValidateToken(r *http.Request) (User, *http.Request, error) {
	tokenAuth := jwtauth.New("HS256", []byte("JWT_SECRET"), nil)
	tokenString := jwtauth.TokenFromHeader(r)

	if tokenString == "" {
		return User{}, r, errors.New("missing token")
	}

	token, err := jwtauth.VerifyToken(tokenAuth, tokenString)
	if err != nil {
		return User{}, r, errors.New("invalid token")
	}

	uuid, ok := token.Get("uuid")
	if !ok {
		return User{}, r, errors.New("missing uuid property in jwt")
	}

	user, err := getUserFromIntegration(r.Header.Get("User-Agent"), tokenString, uuid.(string))
	if err != nil {
		return User{}, r, err
	}

	// Here I can identify my user in any point of my application due the "WithContext" method that include my user information at the request
	// Set the user ID in the request context
	ctx := r.Context()
	ctx = context.WithValue(ctx, struct{}{}, user)
	r = r.WithContext(ctx)

	return user, r, nil
}

// EXAMPLE: Imagine here a function to get a user from integration by token and userAgent
func getUserFromIntegration(userAgent, token, uuid string) (User, error) {
	user := User{
		UUID:  "xpto",
		Email: "John@doe.com",
		Link: struct {
			Roles     []string
			Grantings []string
		}{
			Roles:     []string{"USER"},
			Grantings: []string{},
		},
	}

	return user, nil
}

// DefaultPermissions is a middleware that checks if the user has the default permissions
func DefaultPermissions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, res, err := ValidateToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		r = res

		//TODO: validar grantings
		for _, role := range user.Link.Roles {
			if role == "USER" {
				next.ServeHTTP(w, r)
				return
			}
		}

		//return error forbidden
		http.Error(w, "user does not have enough permissions!", http.StatusForbidden)
	})
}
