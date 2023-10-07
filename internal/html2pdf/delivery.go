package html2pdf

import "github.com/labstack/echo/v4"

// Handlers html2pdf HTTP interface
type Handlers interface {
	HTMLToPDF() echo.HandlerFunc
	URLToPDF() echo.HandlerFunc
}
