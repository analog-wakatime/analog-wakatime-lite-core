package auths

import (
	"encoding/base64"
	"encoding/json"
)

func decodeJWTPart(encoded string, target interface{}) error {
	rawValue, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawValue, target)
}
