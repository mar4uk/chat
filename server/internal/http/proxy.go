package http

import (
	"fmt"
	"net/http"

	"github.com/mar4uk/chat/configs"
	"github.com/mar4uk/chat/internal/app"
)

// Proxy is
type Proxy interface {
	Serve(server configs.ServerConfig) error
}

type proxy struct {
	app app.App
}

// NewProxy is
func NewProxy(app app.App) Proxy {
	return &proxy{
		app: app,
	}
}

func (p *proxy) Serve(server configs.ServerConfig) error {
	r := setupRouter(p.app)
	fmt.Printf("Server is running on http://%s:%s\n", server.Host, server.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", server.Host, server.Port), r); err != nil {
		return err
	}

	return nil
}
