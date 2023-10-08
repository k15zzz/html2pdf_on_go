//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package html2pdf

import (
	"github.com/k15zzz/html2pdf_on_go/internal/models"
	"github.com/labstack/echo/v4"
)

// UseCase Html2pdf use case
type UseCase interface {
	HTMLToPDF(ctx echo.Context, html *models.HTML) (string, error)
	URLToPDF(ctx echo.Context, html *models.URL) (string, error)
}
