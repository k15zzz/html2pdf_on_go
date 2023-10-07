package server

import (
	"github.com/k15zzz/html2pdf_on_go/docs"
	"github.com/k15zzz/html2pdf_on_go/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"

	echoSwagger "github.com/swaggo/echo-swagger"

	html2pdfUseCase "github.com/k15zzz/html2pdf_on_go/internal/html2pdf/usecase"

	html2pdfHttp "github.com/k15zzz/html2pdf_on_go/internal/html2pdf/delivery/http"
)

// MapHandlers Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	docs.SwaggerInfo.Title = "Go example API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	newsUC := html2pdfUseCase.NewHtml2pdfUC(s.cfg, s.logger)

	newsHandlers := html2pdfHttp.NewHtml2pdfHandlers(s.cfg, newsUC, s.logger)

	v1 := e.Group("/api/v1/")

	html2pdfGroup := v1.Group("converts")
	html2pdfHttp.MapNewsRoutes(html2pdfGroup, newsHandlers)

	e.GET("/", func(c echo.Context) error {
		s.logger.Infof("Health check RequestID: %s", utils.GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
