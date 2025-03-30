package main // Entry point of the application

import (
	"github.com/sklevenz/cf-api-server/internal/server"
)

func main() {
	server.StartServer(8080)
}
