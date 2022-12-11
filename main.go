package main

import (
	"log"
	"net/http"
	"wg-go-http/api"
	"wg-go-http/config"

	"gopkg.in/macaron.v1"
)

func main() {
	config := config.ReadConfig()
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(api.AuthMiddleware(config.Secret))
	api.InitApiRoutes(m)

	log.Println(http.ListenAndServe(config.HttpHost, m))
}
