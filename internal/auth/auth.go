package auth

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
)

// CreateJWT returns a JWT given a valid userid+password
func CreateJWT(username string) ([]byte, error) {
	var err error

	signingKey := os.Getenv("SIGNING_PRIVATE_KEY")
	if signingKey == "" {
		log.Fatalln("Unable to load SIGNING_PRIVATE_KEY")
	}
	bytes := []byte(signingKey)
	rsaPrivateKey, err := crypto.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		return nil, err
	}

	jwtExpiresSeconds, err := strconv.ParseInt(
		os.Getenv("JWT_EXPIRES_SECONDS"),
		10,
		32)
	if err != nil {
		return nil, err
	}

	claims := jws.Claims{}
	claims.Set("role", "user")
	claims.SetSubject(username)
	claims.SetIssuer("CART")
	claims.SetIssuedAt(time.Now())
	claims.SetExpiration(time.Now().Add(time.Second * time.Duration(jwtExpiresSeconds)))

	return jws.NewJWT(claims, jws.GetSigningMethod("RS256")).
		Serialize(rsaPrivateKey)
}

// ValidateJWT validates  jwt
func ValidateJWT(j jwt.JWT) error {
	var err error

	signingKey := os.Getenv("SIGNING_PUB_KEY")
	if signingKey == "" {
		log.Fatalln("Unable to load SIGNING_PUB_KEY")
	}
	bytes := []byte(signingKey)
	rsaPublicKey, err := crypto.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		return err
	}
	err = j.Validate(rsaPublicKey, jws.GetSigningMethod("RS256"))
	if err != nil {
		return err
	}

	return nil
}
