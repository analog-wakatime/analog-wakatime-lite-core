package auths

import (
	"strconv"
	"time"
)

func GenerateToken(userID uint, email string) (string, time.Time, error) {
	secret, err := getJWTSecret()
	if err != nil {
		return "", time.Time{}, err
	}

	now := time.Now().UTC()
	expiresAt := now.Add(accessTokenTTL)

	headerPart, err := encodeJWTPart(jwtHeader{Alg: "HS256", Typ: "JWT"})
	if err != nil {
		return "", time.Time{}, err
	}

	claimsPart, err := encodeJWTPart(jwtClaims{
		Subject:   strconv.FormatUint(uint64(userID), 10),
		Email:     email,
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt.Unix(),
	})
	if err != nil {
		return "", time.Time{}, err
	}

	unsignedToken := headerPart + "." + claimsPart
	signature := signJWT(unsignedToken, secret)

	return unsignedToken + "." + signature, expiresAt, nil
}
