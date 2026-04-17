package auths

import (
	"crypto/hmac"
	"errors"
	"time"
)

func parseToken(token string) (jwtClaims, error) {
	secret, err := getJWTSecret()
	if err != nil {
		return jwtClaims{}, err
	}

	parts := splitToken(token)
	if len(parts) != 3 {
		return jwtClaims{}, errors.New("invalid token format")
	}

	var header jwtHeader
	if err := decodeJWTPart(parts[0], &header); err != nil {
		return jwtClaims{}, errors.New("invalid token header")
	}

	if header.Alg != "HS256" || header.Typ != "JWT" {
		return jwtClaims{}, errors.New("unsupported token format")
	}

	expectedSignature := signJWT(parts[0]+"."+parts[1], secret)
	if !hmac.Equal([]byte(expectedSignature), []byte(parts[2])) {
		return jwtClaims{}, errors.New("invalid token signature")
	}

	var claims jwtClaims
	if err := decodeJWTPart(parts[1], &claims); err != nil {
		return jwtClaims{}, errors.New("invalid token payload")
	}

	if claims.Subject == "" {
		return jwtClaims{}, errors.New("token subject is missing")
	}

	if time.Now().UTC().Unix() >= claims.ExpiresAt {
		return jwtClaims{}, errors.New("token expired")
	}

	return claims, nil
}
