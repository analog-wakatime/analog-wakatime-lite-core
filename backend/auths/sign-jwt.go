package auths

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func signJWT(unsignedToken string, secret []byte) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(unsignedToken))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
