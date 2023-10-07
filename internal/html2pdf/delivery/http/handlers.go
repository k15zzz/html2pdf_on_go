package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/k15zzz/html2pdf_on_go/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/k15zzz/html2pdf_on_go/configs"
	"github.com/k15zzz/html2pdf_on_go/internal/html2pdf"
	"github.com/k15zzz/html2pdf_on_go/pkg/httpErrors"
	"github.com/k15zzz/html2pdf_on_go/pkg/logger"
	"github.com/k15zzz/html2pdf_on_go/pkg/utils"
)

var (
	validate *validator.Validate
)

// html2pdfHandlers html2pdf handlers
type html2pdfHandlers struct {
	cfg        *configs.Config
	html2pdfUC html2pdf.UseCase
	logger     logger.Logger
}

// NewHtml2pdfHandlers html to pdf handlers constructor
func NewHtml2pdfHandlers(cfg *configs.Config, html2pdfUC html2pdf.UseCase, logger logger.Logger) html2pdf.Handlers {
	return &html2pdfHandlers{cfg: cfg, html2pdfUC: html2pdfUC, logger: logger}
}

// HTMLToPDF html to pdf handlers generate
func (h html2pdfHandlers) HTMLToPDF() echo.HandlerFunc {
	return func(c echo.Context) error {
		html := new(models.HTML)
		validate = validator.New()

		if err := c.Bind(html); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err := validate.Struct(html); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		pdf := generatePDF(h.html2pdfUC.HTMLToPDF(c, html))

		return c.JSON(http.StatusCreated, pdf)
	}
}

// URLToPDF utl to pdf handlers generate
func (h html2pdfHandlers) URLToPDF() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := new(models.URL)
		validate = validator.New()

		if err := c.Bind(url); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err := validate.Struct(url); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		pdf := generatePDF(h.html2pdfUC.URLToPDF(c, url))

		return c.JSON(http.StatusCreated, pdf)
	}
}

func generatePDF(pdf string) models.PDF {
	return models.PDF{Content: pdf}
}
