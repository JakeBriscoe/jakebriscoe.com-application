package main

import (
	"content-service/internal/api"
	"content-service/internal/database"
	"content-service/internal/spotify"
)

func main() {
	database.ConnectDB()
	spotify.InitializeContent()
	api.StartServer()
}
