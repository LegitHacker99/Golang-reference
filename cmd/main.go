package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"chap1/internal/app"
	db "chap1/internal/database"
)

func main() {
	// starting DB connection
	db.GormInit()
	defer db.GormClose()

	// starting server
	err := app.App()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}
