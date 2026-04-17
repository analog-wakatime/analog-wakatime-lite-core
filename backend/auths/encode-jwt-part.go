package auths

import (
	"encoding/base64"
	"encoding/json"
)

func encodeJWTPart(value interface{}) (string, error) {
	rawValue, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(rawValue), nil
}
