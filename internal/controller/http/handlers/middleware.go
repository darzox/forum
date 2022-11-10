package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"forum/internal/model"
)

type Auth interface {
	SessionCheck(cookie string) (bool, error)
	UserBySession(cookie string) (*model.User, error)
}

type Middleware struct {
	service Auth
}

func CreateMiddleware(service Auth) *Middleware {
	return &Middleware{
		service: service,
	}
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieFromClient := r.Header.Get("Cookie")
		if cookieFromClient == "" {
			next.ServeHTTP(w, r)
			return
		}
		cookieFromClient = strings.ReplaceAll(cookieFromClient, "Session-token=", "")
		ok, err := m.service.SessionCheck(cookieFromClient)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		if !ok {
			next.ServeHTTP(w, r)
			return
		}
		// Refresh cookie expire time after cookie has found
		cookieExpiresAt := time.Now().Add(600 * time.Second)
		http.SetCookie(w, &http.Cookie{
			Name:    "Session-token",
			Value:   cookieFromClient,
			Expires: cookieExpiresAt,
		})
		user, err := m.service.UserBySession(cookieFromClient)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "authorizedUser", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
