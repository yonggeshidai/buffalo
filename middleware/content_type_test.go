package middleware_test

import (
	"testing"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/httptest"
	"github.com/gobuffalo/x/httpx"
	"github.com/stretchr/testify/require"
)

func ctApp() *buffalo.App {
	h := func(c buffalo.Context) error {
		return c.Render(200, render.String(httpx.ContentType(c.Request())))
	}
	a := buffalo.New(buffalo.Options{})
	a.GET("/set", middleware.SetContentType("application/json")(h))
	a.GET("/add", middleware.AddContentType("application/json")(h))
	return a
}

func Test_SetContentType(t *testing.T) {
	r := require.New(t)

	w := httptest.New(ctApp())
	res := w.HTML("/set").Get()
	r.Equal("application/json", res.Body.String())

	req := w.HTML("/set")
	req.Headers["Content-Type"] = "text/plain"

	res = req.Get()
	r.Equal("application/json", res.Body.String())
}

func Test_AddContentType(t *testing.T) {
	r := require.New(t)

	w := httptest.New(ctApp())
	res := w.HTML("/add").Get()
	r.Equal("application/json", res.Body.String())

	req := w.HTML("/add")
	req.Headers["Content-Type"] = "text/plain"

	res = req.Get()
	r.Equal("text/plain", res.Body.String())
}
