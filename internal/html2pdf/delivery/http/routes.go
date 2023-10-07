package http

import (
	"github.com/k15zzz/html2pdf_on_go/internal/html2pdf"
	"github.com/labstack/echo/v4"
)

// MapNewsRoutes Map news routes
func MapNewsRoutes(group *echo.Group, h html2pdf.Handlers) {
	group.POST("/html-to-pdf", h.HTMLToPDF())
	group.POST("/url-to-pdf", h.URLToPDF())
}
