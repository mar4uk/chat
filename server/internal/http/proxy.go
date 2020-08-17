package http

import (
	"fmt"
	"net/http"

	"github.com/mar4uk/chat/configs"
	"github.com/mar4uk/chat/internal/app"
	"github.com/mar4uk/chat/internal/auth"
)

// Proxy is
type Proxy interface {
	Serve() error
}

type proxy struct {
	app  app.App
	auth auth.Auth
	host string
	port uint16
}

// NewProxy is
func NewProxy(app app.App, auth auth.Auth, config configs.ServerConfig) Proxy {
	return &proxy{
		app:  app,
		auth: auth,
		host: config.Host,
		port: config.Port,
	}
}

func (p *proxy) Serve() error {
	r := setupRouter(p.app, p.auth)
	addr := fmt.Sprintf("%s:%d", p.host, p.port)
	fmt.Printf("Server is running on http://%s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		return err
	}

	return nil
}
