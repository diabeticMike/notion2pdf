package notion2pdf

import (
	"encoding/json"
	"github.com/diabeticMike/notion2pdf/entity"
	"github.com/diabeticMike/notion2pdf/pdf"
)

func convert(body []byte) (*pdf.File, error) {
	var page entity.Result
	err := json.Unmarshal(body, &page)
	if err != nil {
		return nil, err
	}

	f := pdf.NewFile()
	for i := range page.Blocks {
		switch page.Blocks[i].Type {
		case "paragraph":
			if page.Blocks[i].Paragraph == nil {
				continue
			}
			for j := range page.Blocks[i].Paragraph.RichText {
				f.Paragraph(page.Blocks[i].Paragraph.RichText[j].Text)
			}
		case "quote":
			if page.Blocks[i].Quote == nil {
				continue
			}
			for j := range page.Blocks[i].Quote.RichText {
				f.Quote(page.Blocks[i].Quote.RichText[j].Text.Content)
			}
		case "heading_1":
			if page.Blocks[i].Heading1 == nil {
				continue
			}
			for j := range page.Blocks[i].Heading1.RichText {
				f.Heading1(page.Blocks[i].Heading1.RichText[j].Text.Content)
			}
		case "heading_2":
			if page.Blocks[i].Heading2 == nil {
				continue
			}
			for j := range page.Blocks[i].Heading2.RichText {
				f.Heading2(page.Blocks[i].Heading2.RichText[j].Text.Content)
			}
		case "heading_3":
			if page.Blocks[i].Heading3 == nil {
				continue
			}
			for j := range page.Blocks[i].Heading3.RichText {
				f.Heading3(page.Blocks[i].Heading3.RichText[j].Text.Content)
			}
		case "bookmark":
			if page.Blocks[i].Bookmark == nil {
				continue
			}
			f.Bookmark(page.Blocks[i].Bookmark.URL)
		case "callout":
			if page.Blocks[i].Callout == nil {
				continue
			}
			for j := range page.Blocks[i].Callout.RichText {
				f.Heading3(page.Blocks[i].Callout.RichText[j].Text.Content)
			}
		case "numbered_list_item":
			if page.Blocks[i].NumberedListItem == nil {
				continue
			}
			for j := range page.Blocks[i].NumberedListItem.RichText {
				f.NumberedListItem(page.Blocks[i].NumberedListItem.RichText[j].Text)
			}
			if len(page.Blocks) > i+1 && page.Blocks[i+1].Type != "numbered_list_item" {
				f.EndNumberedList()
			}
		case "bulleted_list_item":
			if page.Blocks[i].BulledListItem == nil {
				continue
			}
			for j := range page.Blocks[i].BulledListItem.RichText {
				f.BulledListItem(page.Blocks[i].BulledListItem.RichText[j].Text)
			}
		}
	}

	return f, nil
}

func ConvertAndSaveToFile(body []byte, filepath string) error {
	f, err := convert(body)
	if err != nil {
		return err
	}

	return f.SaveToFile(filepath)
}

func ConvertAndSave(body []byte) ([]byte, error) {
	f, err := convert(body)
	if err != nil {
		return nil, err
	}

	pdfBytes, err := f.Save()
	if err != nil {
		return nil, err
	}
	return pdfBytes, nil
}
