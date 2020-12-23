package http

import (
	"fmt"
	"net/http"

	"github.com/mar4uk/chat/configs"
	"github.com/mar4uk/chat/internal/app"
	"github.com/mar4uk/chat/internal/auth"
	"github.com/mar4uk/chat/internal/logger"
)

// Proxy is
type Proxy interface {
	Serve() error
}

type proxy struct {
	app    app.App
	auth   auth.Auth
	host   string
	port   uint16
	logger *logger.Logger
}

// NewProxy is
func NewProxy(app app.App, auth auth.Auth, config configs.ServerConfig, logger *logger.Logger) Proxy {
	return &proxy{
		app:    app,
		auth:   auth,
		host:   config.Host,
		port:   config.Port,
		logger: logger,
	}
}

func (p *proxy) Serve() error {
	r := setupRouter(p.app, p.auth, p.logger)
	addr := fmt.Sprintf("%s:%d", p.host, p.port)
	fmt.Printf("Server is running on http://%s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		return err
	}

	return nil
}
