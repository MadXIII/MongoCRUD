package main

import (
	"context"
	"log"
	"time"

	"github.com/madxiii/mongocrud/internal/database"
	"github.com/madxiii/mongocrud/internal/server"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	client, err := database.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Fatal main error: %s", err)
	}

	defer client.Disconnect(ctx)

	if err = client.Connect(ctx); err != nil {
		log.Fatalf("Fatal main Connect error: %s", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("Fatal main Ping error: %s", err)
	}

	store := database.Store{}
	store.InitCollection(client)

	serv := server.Constr(store)

	if err = serv.Router.Run(":8989"); err != nil {
		log.Fatal(err.Error())
	}
}
