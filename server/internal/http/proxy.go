package http

import (
	"fmt"
	"net/http"

	"github.com/mar4uk/chat/internal/app"
)

// Proxy is
type Proxy interface {
	Serve() error
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

func (p *proxy) Serve() error {
	r := setupRouter(p.app)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}

	return nil
}
