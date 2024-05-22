package main

import (
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"os"
)

func main() {
	// Create a new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	// Create a new page
	page := wkhtmltopdf.NewPage("page-1.html")
	page.EnableLocalFileAccess.Set(true)
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	page.Allow.Set(workingDir)
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./statement.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
