package main

import (
	"os"

	"github.com/Daizaikun/Go-Fiber/application"
	"github.com/Daizaikun/Go-Fiber/database"
)

func main() {

	Database := database.ClientMongo()

	app := application.App(Database)

	app.Listen(port("3000"))

}	

func port(p string) string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = p
	}
	return ":" + port
}
