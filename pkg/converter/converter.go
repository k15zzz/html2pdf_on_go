package converter

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"strings"
)

func GeneratorHTMLToPDF(html string) (string, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println(err)
		return "", err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))

	err = pdfg.Create()
	if err != nil {
		log.Println(err)
		return "", err
	}

	//err = pdfg.WriteFile("./simple.pdf")
	//if err != nil {
	//	log.Println(err)
	//	return "", err
	//}

	return pdfg.Buffer().String(), nil
}

func GeneratorURLToPDF(url string) (string, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	page := wkhtmltopdf.NewPage(url)

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	//err = pdfg.WriteFile("./simple.pdf")
	//if err != nil {
	//	log.Println(err)
	//	return "", err
	//}

	return pdfg.Buffer().String(), nil
}
