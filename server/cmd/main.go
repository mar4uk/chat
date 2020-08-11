package main

import (
	"context"
	"log"

	"github.com/mar4uk/chat/internal/app"
	"github.com/mar4uk/chat/internal/http"
	"github.com/mar4uk/chat/internal/store"
)

func main() {
	ctx := context.Background()

	var (
		db    store.Database
		chat  app.App
		proxy http.Proxy
		err   error
	)

	if db, err = store.NewMongoDatabase(ctx, "mongodb://root:password@localhost:27017", "chat"); err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close(ctx)

	chat = app.NewApp(db)

	proxy = http.NewProxy(chat)

	if err = proxy.Serve(); err != nil {
		log.Fatal(err)
		return
	}
}
