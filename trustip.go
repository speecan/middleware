package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/speecan/util/trustip"
)

// TrustIPConfig for auth config
type TrustIPConfig struct {
	Skipper  middleware.Skipper
	TrustIPs []string
}

var (
	// DefaultTrustIPConfig is used by default
	DefaultTrustIPConfig = TrustIPConfig{
		Skipper: middleware.DefaultSkipper,
	}
)

// TrustIP default auth func
func TrustIP(trustList []string) echo.MiddlewareFunc {
	c := DefaultTrustIPConfig
	c.TrustIPs = trustList
	return TrustIPWithConfig(c)
}

// TrustIPWithConfig needs config
func TrustIPWithConfig(config TrustIPConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			if trustip.IsContained(c.RealIP(), config.TrustIPs) {
				return next(c)
			}
			return c.NoContent(http.StatusForbidden)
		}
	}
}
