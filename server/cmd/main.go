package main

import (
	"log"

	"github.com/naufalhakm/go-chat/database"
	"github.com/naufalhakm/go-chat/internal/user"
	"github.com/naufalhakm/go-chat/router"
)

func main() {
	dbConn, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
