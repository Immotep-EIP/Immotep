package pdf

import (
	"bytes"
	"io"

	"github.com/go-pdf/fpdf"
	"keyz/backend/prisma/db"
	"keyz/backend/utils"
)

var Test = false

type Header float64

const (
	H1 Header = 20.0
	H2 Header = 18.0
	H3 Header = 16.0
	H4 Header = 14.0
)

type CaptureWriter struct {
	io.Writer
	buffer bytes.Buffer
}

func (cw *CaptureWriter) Write(p []byte) (int, error) {
	return cw.buffer.Write(p)
}

func (cw *CaptureWriter) GetData() []byte {
	return cw.buffer.Bytes()
}

type PDF struct {
	pdf *fpdf.Fpdf
}

func NewPDF() PDF {
	res := PDF{
		pdf: fpdf.New("P", "mm", "A4", ""),
	}

	basePath := utils.Ternary(Test, "", "services/pdf/")
	res.pdf.AddUTF8Font("roboto", "", basePath+"Roboto-Regular.ttf")
	res.pdf.AddUTF8Font("roboto", "B", basePath+"Roboto-Bold.ttf")
	res.pdf.AddUTF8Font("roboto", "I", basePath+"Roboto-Italic.ttf")
	res.pdf.AddUTF8Font("roboto", "BI", basePath+"Roboto-BoldItalic.ttf")
	res.pdf.SetFont("roboto", "", 12)
	res.pdf.AddPage()
	return res
}

func (irp *PDF) Save(name string) error {
	return irp.pdf.OutputFileAndClose(name)
}

func (irp *PDF) Output() ([]byte, error) {
	w := &CaptureWriter{}
	err := irp.pdf.Output(w)
	if err != nil {
		return nil, err
	}
	return w.GetData(), nil
}

func (irp *PDF) GetWidth() float64 {
	docW, _ := irp.pdf.GetPageSize()
	marginL, _, marginR, _ := irp.pdf.GetMargins()
	return docW - marginL - marginR
}

func (irp *PDF) Ln(n float64) {
	irp.pdf.Ln(n)
}

func (irp *PDF) AddCenteredTitle(title string, header Header) {
	irp.pdf.SetFont("roboto", "B", float64(header))
	irp.pdf.CellFormat(irp.GetWidth(), 10, title, "", 0, "C", false, 0, "")
	irp.pdf.Ln(10)
}

func (irp *PDF) AddTitle(title string, header Header) {
	irp.pdf.SetFont("roboto", "B", float64(header))
	irp.pdf.Cell(irp.GetWidth(), 10, title)
	irp.pdf.Ln(10)
}

func (irp *PDF) AddText(text string) {
	irp.pdf.SetFont("roboto", "", 12)
	irp.pdf.Cell(irp.GetWidth(), 10, text)
	irp.pdf.Ln(5)
}

func (irp *PDF) Add2Texts(text1, text2 string) {
	irp.pdf.SetFont("roboto", "", 12)
	irp.pdf.Cell(irp.GetWidth()/2, 10, text1)
	irp.pdf.Cell(irp.GetWidth()/2, 10, text2)
	irp.pdf.Ln(5)
}

func (irp *PDF) AddMultiLineText(text string) {
	irp.pdf.Ln(5)
	irp.pdf.SetFont("roboto", "", 12)
	irp.pdf.MultiCell(irp.GetWidth(), 5, text, "", "L", false)
}

func (irp *PDF) AddLine() {
	docW, _ := irp.pdf.GetPageSize()
	marginL, _, marginR, _ := irp.pdf.GetMargins()
	irp.pdf.Line(marginL, irp.pdf.GetY(), docW-marginR, irp.pdf.GetY())
}

func (irp *PDF) AddImages(images []db.ImageModel) {
	docW, docH := irp.pdf.GetPageSize()
	marginL, _, marginR, marginB := irp.pdf.GetMargins()

	maxWidth := docW - marginR
	imageHeight := 69.35
	currentX := irp.pdf.GetX()
	currentY := irp.pdf.GetY()
	for _, picture := range images {
		imageOptions := fpdf.ImageOptions{
			ReadDpi:   true,
			ImageType: string(picture.Type),
		}
		info := irp.pdf.RegisterImageOptionsReader(picture.ID, imageOptions, bytes.NewReader(picture.Data))
		imageWidth := info.Width() * imageHeight / info.Height()
		if currentX+imageWidth > maxWidth {
			currentX = marginL
			currentY += imageHeight + 5
		}
		if currentY+imageHeight > docH-marginB {
			irp.pdf.AddPage()
			currentY = irp.pdf.GetY()
		}
		irp.pdf.ImageOptions(picture.ID, currentX, currentY, 0, imageHeight, false, imageOptions, 0, "")
		currentX += imageWidth + 5
	}
	irp.pdf.SetY(currentY + imageHeight + 5)
}
