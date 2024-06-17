package media

import (
	"log"
	"monify/lib/auth"
	"net/http"
)

func Start(config Config) {

	authMiddleware := auth.AuthMiddleware{JwtSecret: config.JwtSecret}
	http.Handle("/image", authMiddleware.HttpMiddleware(http.HandlerFunc(uploadImage)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
