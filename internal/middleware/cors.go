package middleware

import (
	"github.com/go-chi/cors"
	appConfig "github.com/sonochiwa/wb-level-0/config"
)

var cfg = appConfig.GetConfig()

func GetCors() cors.Options {
	return cors.Options{
		AllowedOrigins:   cfg.Cors.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: cfg.Cors.AllowCredentials,
		MaxAge:           cfg.Cors.MaxAge,
	}
}
