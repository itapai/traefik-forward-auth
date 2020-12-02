package server

import (
	"bytes"
	"fmt"
	"log"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func ValidateToken(jwks string, token []byte) (jwt.Token, error) {
	kset, err := jwk.FetchHTTP(jwks)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to parse JWK: %v", err)
	}
	pubkey := kset.Keys[0].X509CertChain()[0].PublicKey

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
