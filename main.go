package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

const (
	TOKEN         = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICI3VThTSjU2bWdDTEZTOVhHTTFXNTU2SVVFLVZXdUlsYzF5bF9WU3lPaFFVIn0.eyJleHAiOjE2MDY5Mjg3MTIsImlhdCI6MTYwNjkyODQxMiwianRpIjoiM2M2OTlkNTUtMGNiNy00NzZkLWFjM2MtYjZmZjJhMjU5ZDY2IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL2F1dGgvcmVhbG1zL2RlbW8iLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiYzM2NjU3YzgtNTQ1ZS00NWFmLTkxMTEtOGMwZjMxN2ZjYTVkIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiZmF1dGgiLCJzZXNzaW9uX3N0YXRlIjoiOWZlMTczNzAtYjBmMC00MzgzLThkMDAtODY3OWU2MDdmNDE3IiwiYWNyIjoiMSIsInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsIm5hbWUiOiJEZSBNbyIsInByZWZlcnJlZF91c2VybmFtZSI6ImRlbW8iLCJnaXZlbl9uYW1lIjoiRGUiLCJmYW1pbHlfbmFtZSI6Ik1vIiwiZW1haWwiOiJkQGUubW8ifQ.VFGRl1xyDpMOjfNVwJmdNx5wvTB514rIy_5AQYla2TOQ5DpyjPEXIgdtbVmtoIlhtXTS2qtG_b2CUSuAyIjZ2fjNyumVvaN4ncqcSSt1O6sUMLDqWU1UDPWOct4V0NvXi1j-WIOll7r7mj8Fm9M0KxUYoAigJVSeIiyi35awfkVdkD92U0tkP2Syny9I89yULVYgexkuyXW5ti3hhgqdw97Kj3dWVVcNB16UoKHFe8siEbWAm1KSGMDlLb0c-jv0-9EYrqci5y28GI1_Fj4TLwdNCWZ_QcHSRkwmcwTnwOg-ENdqbucvN18uulYVhW-r0268IUBW8S3qM9-KYc-vLQ"
	INVALID_TOKEN = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.POstGetfAytaZS82wHcjoTyoqhMyxXiWdR7Nn7A29DNSl0EiXLdwJ6xC6AfgZWF1bOsS_TuYI3OG85AmiExREkrS6tDfTQ2B3WXlrr-wp5AokiRbz3_oB4OxG-W9KcEEbDRcZc0nH3L7LzYptiy1PtAylQGxHTWZXtGz4ht0bAecBgmpdgXMguEIcoqPJ1n3pIWk_dUZegpqx0Lka21H6XxUTxiy8OcaarA8zdnPUnV6AmNP3ecFawIFYdvJB_cm-GvpCSbr8G8y_Mllj8f4x9nBH8pQux89_6gUY618iYv7tuPWBFfEbLxtF2pZS6YC1aSfLQxeNe8djT9YjpvRZA"
)

func main() {
	log.Fatal(isTokenValid([]byte(TOKEN)))
}

func isTokenValid(token []byte) error {
	u := "http://keycloak.localhost/auth/realms/demo/protocol/openid-connect/certs"
	jwks, err := jwk.FetchHTTP(u)
	if err != nil {
		return fmt.Errorf("failed to parse JWK: %v", err)
	}
	pubkey := jwks.Keys[0].X509CertChain()[0].PublicKey

	t, err := jwt.Parse(bytes.NewReader(token), jwt.WithVerify(jwa.RS256, pubkey))
	if err != nil {
		return fmt.Errorf("failed to parse JWT token: %v", err)
	}

	// buf, err := json.MarshalIndent(token, "", "  ")
	// if err != nil {
	// 	return fmt.Errorf("failed to generate JSON: %v", err)
	// }
	// fmt.Printf("%s\n", buf)

	// check if token is valid
	return jwt.Verify(t)
}
