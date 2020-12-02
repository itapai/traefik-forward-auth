package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
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

	_, err := s.verifyToken([]byte(bearerToken))
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

func (s *server) verifyToken(token []byte) (jwt.Token, error) {
	jwks, err := jwk.FetchHTTP(s.c.JWKS)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to parse JWK: %v", err)
	}
	pubkey := jwks.Keys[0].X509CertChain()[0].PublicKey

	t, err := jwt.Parse(bytes.NewReader(token), jwt.WithVerify(jwa.RS256, pubkey))
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT token: %v", err)
	}

	// buf, err := json.MarshalIndent(token, "", "  ")
	// if err != nil {
	// 	return fmt.Errorf("failed to generate JSON: %v", err)
	// }
	// fmt.Printf("%s\n", buf)

	// check if token is valid
	err = jwt.Verify(t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
