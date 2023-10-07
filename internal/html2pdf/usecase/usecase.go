package usecase

import (
	"github.com/k15zzz/html2pdf_on_go/configs"
	"github.com/k15zzz/html2pdf_on_go/internal/html2pdf"
	"github.com/k15zzz/html2pdf_on_go/internal/models"
	"github.com/k15zzz/html2pdf_on_go/pkg/converter"
	"github.com/k15zzz/html2pdf_on_go/pkg/logger"
	"github.com/labstack/echo/v4"
)

const (
	basePrefix    = "api-news:"
	cacheDuration = 3600
)

// html2pdfUC News UseCase
type html2pdfUC struct {
	cfg    *configs.Config
	logger logger.Logger
}

// NewHtml2pdfUC html to pdf UseCase constructor
func NewHtml2pdfUC(cfg *configs.Config, logger logger.Logger) html2pdf.UseCase {
	return &html2pdfUC{cfg: cfg, logger: logger}
}

// HTMLToPDF convert html (string) to PDF(string)
func (u *html2pdfUC) HTMLToPDF(ctx echo.Context, html *models.HTML) string {
	return converter.GeneratorHTMLToPDF(html.Content)
}

// URLToPDF convert html (string) to PDF(string)
func (u *html2pdfUC) URLToPDF(ctx echo.Context, url *models.URL) string {
	return converter.GeneratorURLToPDF(url.Link)
}
