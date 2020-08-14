package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mar4uk/chat/configs"

	"github.com/mar4uk/chat/internal/app"
	"github.com/mar4uk/chat/internal/http"
	"github.com/mar4uk/chat/internal/store"
)

func main() {
	ctx := context.Background()

	var (
		config *configs.Config
		db     store.Database
		chat   app.App
		proxy  http.Proxy
		err    error
	)

	if config, err = configs.ParseConfig("configs/config.yml"); err != nil {
		log.Fatal(err)
		return
	}

	addr := fmt.Sprintf("%s://%s:%s@%s:%s", config.Database.Protocol,
		config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port)

	if db, err = store.NewMongoDatabase(ctx, addr, config.Database.Name); err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close(ctx)

	chat = app.NewApp(db)

	proxy = http.NewProxy(chat)

	if err = proxy.Serve(config.Server); err != nil {
		log.Fatal(err)
		return
	}
}
