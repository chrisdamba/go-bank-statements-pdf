package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"html/template"
	"log"
	"os"
)

type StatementData struct {
	AccountHolder string
	DateCreated   string
	Transactions  []Transaction
}

type Transaction struct {
	Date            string
	MainDescription string
	SubDescription  string
	Out             float64
	In              float64
	Balance         float64
}

func main() {
	// Start a headless Chrome instance
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Prepare data for the template
	data := StatementData{
		AccountHolder: "Sandra Saulgrieze",
		DateCreated:   "20 May 2023",
		Transactions: []Transaction{
			{
				Date:            "3 Feb 2023",
				MainDescription: "Apple Pay Top-Up by *5453",
				SubDescription:  "",
				Out:             0,
				In:              50.00,
				Balance:         52.52,
			},
			{
				Date:            "3 Feb 2023",
				MainDescription: "Apple Pay Top-Up by *5453",
				SubDescription:  "",
				Out:             0,
				In:              100.00,
				Balance:         152.52,
			},
			{
				Date:            "3 Feb 2023",
				MainDescription: "To LINA MILLER SAULGRIEZE",
				SubDescription:  "To: LINA MILLER SAULGRIEZE",
				Out:             100.00,
				In:              0,
				Balance:         52.52,
			},
			{
				Date:            "7 Feb 2023",
				MainDescription: "To LINA MILLER SAULGRIEZE",
				SubDescription:  "To: LINA MILLER SAULGRIEZE",
				Out:             10.00,
				In:              0,
				Balance:         42.52,
			},
		},
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}

	// Render the template to a buffer
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}

	// Navigate to the HTML content
	err = chromedp.Run(ctx,
		chromedp.Navigate("data:text/html,"+buf.String()),
	)
	if err != nil {
		log.Fatal(err)
	}
	var pageBuf []byte
	// Generate the PDF
	err = chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pageBuf, _, err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			return err
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Save the PDF
	err = os.WriteFile(
		"statement.pdf",
		pageBuf,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

/*
package main

import (
	"context"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"io/ioutil"
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
*/
