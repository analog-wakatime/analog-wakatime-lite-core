package auths

import (
	"analog-wakatime-lite-core/config"
	"errors"
)

func getJWTSecret() ([]byte, error) {
	secret := config.ConfigGetJWTSecret()
	if secret == "" {
		return nil, errors.New("JWT_SECRET is not configured")
	}

	return []byte(secret), nil
}
