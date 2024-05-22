## HTML to PDF Converter (Go)

This simple Go application converts an HTML file (`page-1.html`) into a PDF document (`statement.pdf`) using the `wkhtmltopdf` tool.

### Prerequisites

*   **Go:** Ensure you have Go installed on your system. You can download it from the official website: [https://go.dev/](https://go.dev/)
*   **wkhtmltopdf:** Install wkhtmltopdf. It's a command-line tool that renders HTML into PDF and various image formats. Download and install it from: [https://wkhtmltopdf.org/](https://wkhtmltopdf.org/)
*   **go-wkhtmltopdf library:** This Go library provides a wrapper around wkhtmltopdf, making it easier to use within your Go code. Install it using:
    ```bash
    go get github.com/SebastiaanKlippert/go-wkhtmltopdf
    ```

### Usage

1.  **Place your HTML file:** Make sure your `page-1.html` file is in the same directory as the `main.go` file.

2.  **Build:** In your terminal, navigate to the directory and run:
    ```bash
    make build
    ```
    This will create the `converter` executable.
3.  **Run:** To convert your HTML to PDF, run:
    ```bash
    make run   # or simply 'make'
    ```
    This will execute the `converter` executable and generate the `statement.pdf` file.
4.  **Clean:** To clean up the generated files, run:
    ```bash
    make clean
    ```

### Customisation

*   **Output file name:** Change the `"./statement.pdf"` argument in `pdfg.WriteFile()` to your desired output file name.
*   **wkhtmltopdf options:** The `pdfg` object provides various options to customize the PDF generation process (e.g., DPI, orientation, margins). Refer to the `go-wkhtmltopdf` documentation for available options.
*   **Input file:** Modify the `page := wkhtmltopdf.NewPage("page-1.html")` line to specify a different HTML file if needed.

### How it Works

The code does the following:

1.  Creates a new PDF generator using `wkhtmltopdf.NewPDFGenerator()`.
2.  Sets global options like DPI and orientation.
3.  Creates a new page object from your HTML file using `wkhtmltopdf.NewPage("page-1.html")`.
4.  Adds the page to the PDF generator.
5.  Generates the PDF document in memory.
6.  Writes the PDF content to a file on disk.

### Troubleshooting

*   **"wkhtmltopdf not found":** Make sure wkhtmltopdf is installed and in your system's PATH.
*   **"Failed to load about:blank":** This error usually indicates that wkhtmltopdf cannot access your HTML file or its resources (CSS, images). Ensure they are in the correct location and have the correct permissions.

### Example `page-1.html`

```html
<!DOCTYPE html>
<html>
<head>
    <title>Example Page</title>
    <link rel="stylesheet" href="assets/styles.css">
</head>
<body>
    <h1>Hello, World!</h1>
    <p>This is an example HTML page to be converted to PDF.</p>
</body>
</html>
```

**Note:** Make sure that your stylesheet (styles.css) and other assets are in the same assets directory as your HTML file.
