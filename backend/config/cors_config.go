package config

import (
	"net/url"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
)

var allowedOrigins = []string{
	"http://localhost:3000",
}

func normalizeOrigin(origin string) string {
	origin = strings.TrimSpace(origin)
	origin = strings.TrimSuffix(origin, "/")
	if origin == "" {
		return ""
	}

	parsedOrigin, err := url.Parse(origin)
	if err != nil || parsedOrigin.Scheme == "" || parsedOrigin.Host == "" {
		return strings.ToLower(origin)
	}

	scheme := strings.ToLower(parsedOrigin.Scheme)
	host := strings.ToLower(parsedOrigin.Hostname())
	port := parsedOrigin.Port()

	if (scheme == "http" && port == "80") || (scheme == "https" && port == "443") {
		port = ""
	}

	if port != "" {
		return scheme + "://" + host + ":" + port
	}
	return scheme + "://" + host
}

func IsAllowedOrigin(origin string) bool {
	originClean := normalizeOrigin(origin)
	for _, allowedOrigin := range allowedOrigins {
		if originClean == normalizeOrigin(allowedOrigin) {
			return true
		}
	}
	return false
}

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOriginFunc:  IsAllowedOrigin,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "ngrok-skip-browser-warning", "X-Requested-With", "Access-Control-Request-Method", "Access-Control-Request-Headers", "id", "Id"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}
