package schedule

import (
	"example.com/go-demo-1/downlader"
	"github.com/signintech/gopdf"
)

//595.28, 841.89 = A4
const (
	FONTURL = "https://github.com/google/fonts/raw/master/ofl/daysone/DaysOne-Regular.ttf"
	PDF     = "-pdf.pdf"
	TTF     = "-font.ttf"
	FONT    = "daysone"
	STYLE   = ""
	SIZE    = 20
	WIDTH   = 595.28
	HEIGHT  = 841.89
	TEMP    = "temp"
)

type Schedule struct {
	Name, Url string
}

func (s Schedule) Process() {
	var err error

	// Download a Font
	fontFile := TEMP + TTF
	pdfFile := TEMP + PDF

	downloadErrors := make(chan error, 2)

	go downlader.DownloadFile(fontFile, FONTURL, downloadErrors)
	go downlader.DownloadFile(pdfFile, s.Url, downloadErrors)

	if err = <-downloadErrors; err != nil {
		panic(err)
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: WIDTH, H: HEIGHT}})
	pdf.AddPage()

	err = pdf.AddTTFFont(FONT, fontFile)
	if err != nil {
		panic(err)
	}

	err = pdf.SetFont(FONT, STYLE, SIZE)
	if err != nil {
		panic(err)
	}

	// Color the page
	pdf.SetLineWidth(0.1)
	pdf.SetFillColor(124, 252, 0) //setup fill color
	pdf.RectFromUpperLeftWithStyle(50, 100, 400, 600, "FD")
	pdf.SetFillColor(0, 0, 0)

	pdf.SetX(50)
	pdf.SetY(50)
	pdf.Cell(nil, "Import existing PDF into GoPDF Document")

	// Import page 1
	tpl1 := pdf.ImportPage(pdfFile, 1, "/MediaBox")

	// Draw pdf onto page
	pdf.UseImportedTemplate(tpl1, 50, 100, 400, 0)

	pdf.WritePdf(s.Name)
}
