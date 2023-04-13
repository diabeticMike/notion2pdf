package pdf

import (
	"bytes"
	"github.com/diabeticMike/notion2pdf/entity"
	"github.com/go-pdf/fpdf"
	"strconv"
)

type File struct {
	pdf       *fpdf.Fpdf
	name      string
	listCount int
}

func NewFile(name string) *File {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	return &File{pdf: pdf, name: name, listCount: 1}
}

func (f *File) Paragraph(text entity.Text) {
	if text.Link != nil {
		f.setFontForUrl()
		f.pdf.WriteLinkString(0, text.Content, text.Link.URL)
		f.pdf.Ln(10)
		f.setBaseFont()
		return
	}
	f.pdf.Cell(0, 0, text.Content)
	f.pdf.Ln(10)
}

func (f *File) Callout(text entity.Text) {
	if text.Link != nil {
		f.setFontForUrl()
		f.pdf.WriteLinkString(0, text.Content, text.Link.URL)
		f.pdf.Ln(10)
		f.setBaseFont()
		return
	}
	f.pdf.Cell(0, 0, text.Content)
	f.pdf.Ln(10)
}

func (f *File) SaveToFile() error {
	return f.pdf.OutputFileAndClose(f.name)
}

func (f *File) Save() ([]byte, error) {
	defer f.pdf.Close()
	var buffer bytes.Buffer
	err := f.pdf.Output(&buffer)
	if err != nil {
		return nil, err
	}

	var body []byte
	_, err = buffer.Read(body)
	return nil, err
}

func (f *File) Quote(text string) {
	f.pdf.Cell(0, 0, "| "+text)
	f.pdf.Ln(10)
}

func (f *File) NumberedListItem(text entity.Text) {
	if text.Link != nil {
		f.setFontForUrl()
		f.pdf.WriteLinkString(0, strconv.Itoa(f.listCount)+". "+text.Content, text.Link.URL)
		f.pdf.Ln(10)
		f.setBaseFont()
		f.listCount++
		return
	}
	f.pdf.Cell(0, 0, strconv.Itoa(f.listCount)+". "+text.Content)
	f.pdf.Ln(10)
	f.listCount++
}

func (f *File) BulledListItem(text entity.Text) {
	if text.Link != nil {
		f.setFontForUrl()
		f.pdf.WriteLinkString(0, "* "+text.Content, text.Link.URL)
		f.pdf.Ln(10)
		f.setBaseFont()
		return
	}
	f.pdf.Cell(0, 0, "* "+text.Content)
	f.pdf.Ln(10)
}

func (f *File) EndNumberedList() {
	f.listCount = 1
}

func (f *File) Heading1(text string) {
	f.pdf.SetFont("Arial", "", 20)
	f.pdf.Cell(0, 0, text)

	f.setBaseFont()
	f.pdf.Ln(10)
}

func (f *File) Heading2(text string) {
	f.pdf.SetFont("Arial", "", 16)
	f.pdf.Cell(0, 0, text)

	f.setBaseFont()
	f.pdf.Ln(10)
}

func (f *File) Heading3(text string) {
	f.pdf.SetFont("Arial", "", 12)
	f.pdf.Cell(0, 0, text)

	f.setBaseFont()
	f.pdf.Ln(10)
}

func (f *File) Bookmark(url string) {
	f.setFontForUrl()
	f.pdf.WriteLinkString(0, url, url)
	f.pdf.Ln(10)
	return
	f.setBaseFont()
	f.pdf.Ln(10)
}

func (f *File) setBaseFont() {
	f.pdf.SetFont("Arial", "", 10)
}

func (f *File) setFontForUrl() {
	f.pdf.SetFont("Arial", "I", 10)
}
