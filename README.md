## Dynamic Bank Statement Generator (Go)

This Go application dynamically generates a PDF bank statement from an HTML template (`template.html`) using the `chromedp` library and headless Chrome. It populates the template with transaction data from a Go struct, allowing for flexible and personalized statement creation.

### Prerequisites

*   **Go:** Ensure you have Go installed on your system. You can download it from the official website: [https://go.dev/](https://go.dev/)
*   **chromedp library:** This library is used to control a headless Chrome or Chromium instance using the DevTools Protocol. Install it using:
    ```bash
    go get -u github.com/chromedp/chromedp
    ```

### Usage

1.  **Prepare your data:** Define a Go struct (e.g., `StatementData`) to hold your bank statement information, including account holder details and transaction data.
2.  **Create an HTML template:** Design your `template.html` file. Use Go's template syntax (`{{ }}`) to insert dynamic data into the HTML structure.
3.  **Place your files:** Ensure your `main.go`, `template.html`, and any associated CSS files (e.g., `styles.css`) are in the same directory.
4.  **Build and Run:** Use the provided Makefile to build and run the application:
    ```bash
    make build  # Compile the Go code
    make run    # Execute the compiled program
    ```
    This will generate a PDF file named `statement.pdf` in the same directory.

### How it Works

1.  **Data Initialization:** The Go code initializes a `StatementData` struct with your bank statement information.
2.  **Template Parsing:** The HTML template (`template.html`) is parsed using Go's `html/template` package.
3.  **Template Execution:** The template is executed with the `StatementData`, dynamically filling in the placeholders with your data.
4.  **Headless Chrome Rendering:** The `chromedp` library controls a headless Chrome instance to:
    *   Navigate to the rendered HTML content.
    *   Generate a PDF from the rendered page, including background styles and images.
5.  **PDF Saving:** The generated PDF is saved to `statement.pdf`.

### Code Structure

*   **`StatementData` struct:** Defines the structure of your bank statement data.
*   **`Transaction` struct:** Represents individual transactions within the statement.
*   **`main` function:**
    *   Initializes the `StatementData`.
    *   Parses and executes the HTML template.
    *   Uses `chromedp` to render and save the PDF.

### Customization

*   **Data:** Tailor the `StatementData` and `Transaction` structs to match your specific bank statement format.
*   **Template:** Customize `template.html` to achieve your desired layout and styling.
*   **Output:** Modify the output file name in `os.WriteFile()`.
*   **Chrome Options:** Explore `chromedp` options for further customization of the PDF generation process (e.g., page size, margins).

### Example Data

```go
data := StatementData{
    AccountHolder: "Sandra Saulgrieze",
    DateCreated:   "20 May 2023",
    Transactions: []Transaction{
        {Date: "3 Feb 2023", MainDescription: "Apple Pay Top-Up by *5453", In: 50.00, Balance: 52.52},
        {Date: "3 Feb 2023", MainDescription: "Apple Pay Top-Up by *5453", In: 100.00, Balance: 152.52},
        {Date: "3 Feb 2023", MainDescription: "To LINA MILLER SAULGRIEZE", Out: 100.00, Balance: 52.52},
        {Date: "7 Feb 2023", MainDescription: "To LINA MILLER SAULGRIEZE", Out: 10.00, Balance: 42.52},
    },
}
```

### Example Template (`template.html`)

```html
<!DOCTYPE html>
<html>
<head>
    <title>Bank Statement</title>
    <link rel="stylesheet" href="styles.css"> 
</head>
<body>
    <h1>Account Holder: {{ .AccountHolder }}</h1>
    </body>
</html>
```