package architecture

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Routes interface {
	Register(e *echo.Echo)
}

type Router struct {
	Port        int
	RouteGroups []Routes
}

func (r *Router) Start() error {
	e := echo.New()
	for _, r := range r.RouteGroups {
		r.Register(e)
	}
	fmt.Printf("\n%d Routes registered\n", len(e.Routes()))
	for _, r := range e.Routes() {
		fmt.Printf("\t%s: %s\n", r.Method, r.Path)
	}
	return e.Start(fmt.Sprintf(":%d", r.Port))
}
