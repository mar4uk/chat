package http

import (
	"fmt"
	"net/http"

	"github.com/mar4uk/chat/configs"
	"github.com/mar4uk/chat/internal/app"
)

// Proxy is
type Proxy interface {
	Serve() error
}

type proxy struct {
	app  app.App
	host string
	port uint16
}

// NewProxy is
func NewProxy(app app.App, config configs.ServerConfig) Proxy {
	return &proxy{
		app:  app,
		host: config.Host,
		port: config.Port,
	}
}

func (p *proxy) Serve() error {
	r := setupRouter(p.app)
	addr := fmt.Sprintf("%s:%d", p.host, p.port)
	fmt.Printf("Server is running on http://%s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		return err
	}

	return nil
}
