package utils

import (
	"context"
	"github.com/k15zzz/html2pdf_on_go/pkg/logger"
	"github.com/labstack/echo/v4"
	"time"
)

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// GetCtxWithReqID Get ctx with timeout and request id from echo context
func GetCtxWithReqID(c echo.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*15)
	ctx = context.WithValue(ctx, ReqIDCtxKey{}, GetRequestID(c))
	return ctx, cancel
}

// GetRequestCtx Get context  with request id
func GetRequestCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// GetRequestID Get request id from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// GetConfigPath Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./configs/docker"
	}
	return "./configs/local"
}

// LogResponseError Error response with logging error for echo context
func LogResponseError(ctx echo.Context, logger logger.Logger, err error) {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(ctx),
		GetIPAddress(ctx),
		err,
	)
}

// GetIPAddress Get user ip address
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}
