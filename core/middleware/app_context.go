package middleware

import (
	"github.com/labstack/echo"
	"github.com/starptech/go-web/context"
)

func AppContext(cc *context.AppContext) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc.Context = c
			return h(cc)
		}
	}
}
