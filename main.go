package main

import (
	"canvas/database"
	"canvas/routes"
)

func main() {

	database.Migrate()
	r := routes.SetupRouter()
	_ = r.Run(":7001")
}
