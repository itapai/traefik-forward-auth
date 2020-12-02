package server

import (
	"net/http"
	"strings"
)

type server struct {
	c *Config
}

func NewServer(c *Config) *server {
	return &server{
		c: c,
	}
}

// JWTHandler checks for an Authorization header, extracts the JWT token and
// returns 200 or 401 if the token is valid or not
func (s *server) JWTHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	bearerToken := authHeader[len("Bearer "):]

	_, err := ValidateToken(s.c.JWKS, []byte(bearerToken))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
		return
	}

	// get claims
	// spew.Dump(token.Get("realm_access"))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func Run(c *Config) error {
	s := NewServer(c)

	http.HandleFunc("/", s.JWTHandler)

	return http.ListenAndServe(c.Addr, nil)
}
