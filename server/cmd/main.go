package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mar4uk/chat/configs"

	"github.com/mar4uk/chat/internal/app"
	"github.com/mar4uk/chat/internal/auth"
	"github.com/mar4uk/chat/internal/http"
	"github.com/mar4uk/chat/internal/logger"
	"github.com/mar4uk/chat/internal/store"
)

func main() {
	ctx := context.Background()

	var (
		config *configs.Config
		db     store.Database
		chat   app.App
		ah     auth.Auth
		proxy  http.Proxy
		err    error
	)

	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
		return
	}

	if config, err = configs.ParseConfig("configs/config.yml"); err != nil {
		logger.Logger.Fatal(err)
		return
	}

	addr := fmt.Sprintf("%s://%s:%s@%s:%s", config.Mongo.Protocol,
		config.Mongo.Username, config.Mongo.Password, config.Mongo.Host, config.Mongo.Port)

	if db, err = store.NewMongoDatabase(ctx, addr, config.Mongo.Name); err != nil {
		logger.Logger.Fatal(err)
		return
	}
	defer db.Close(ctx)

	chat = app.NewApp(db)
	ah = auth.NewAuth(db)

	proxy = http.NewProxy(chat, ah, config.Server, logger)

	if err = proxy.Serve(); err != nil {
		logger.Logger.Fatal(err)
		return
	}
}
