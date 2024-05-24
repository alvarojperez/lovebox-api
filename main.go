package main

import (
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	godotenv.Load()

	openDBConnection()
	defer db.Close()

	server := newApiServer(":" + getEnvVariable("PORT"))
	server.run()
}
