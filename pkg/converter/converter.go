package converter

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"strings"
)

func GeneratorHTMLToPDF(html string) string {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//err = pdfg.WriteFile("./simple.pdf")
	//if err != nil {
	//	log.Fatal(err)
	//}

	return pdfg.Buffer().String()
}

func GeneratorURLToPDF(url string) string {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	page := wkhtmltopdf.NewPage(url)

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//err = pdfg.WriteFile("./simple.pdf")
	//if err != nil {
	//	log.Fatal(err)
	//}

	return pdfg.Buffer().String()
}
