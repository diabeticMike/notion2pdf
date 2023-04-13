package entity

import "time"

type Result struct {
	Object string  `json:"object"`
	Blocks []Block `json:"results"`
}

type Block struct {
	Object           string            `json:"object"`
	ID               string            `json:"id"`
	Parent           Parent            `json:"parent"`
	CreatedTime      time.Time         `json:"created_time"`
	LastEditedTime   time.Time         `json:"last_edited_time"`
	CreatedBy        CreatedBy         `json:"created_by"`
	LastEditedBy     LastEditedBy      `json:"last_edited_by"`
	HasChildren      bool              `json:"has_children"`
	Archived         bool              `json:"archived"`
	Type             string            `json:"type"`
	Paragraph        *Paragraph        `json:"paragraph,omitempty"`
	Quote            *Quote            `json:"quote,omitempty"`
	Heading1         *Heading          `json:"heading_1,omitempty"`
	Heading2         *Heading          `json:"heading_2,omitempty"`
	Heading3         *Heading          `json:"heading_3,omitempty"`
	Bookmark         *Bookmark         `json:"bookmark,omitempty"`
	Callout          *Callout          `json:"callout,omitempty"`
	NumberedListItem *NumberedListItem `json:"numbered_list_item,omitempty"`
	BulledListItem   *BulledListItem   `json:"bulleted_list_item,omitempty"`
}

type Parent struct {
	Type   string `json:"type"`
	PageID string `json:"page_id"`
}

type CreatedBy struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type LastEditedBy struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type RichText struct {
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`
}

type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link"`
}

type NumberedListItem struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color"`
}

type BulledListItem struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color"`
}

type Link struct {
	URL string `json:"url"`
}

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

type Paragraph struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color"`
}

type Heading struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color"`
}

type Quote struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color"`
}

type Callout struct {
	RichText []RichText `json:"rich_text"`
	Color    string     `json:"color"`
}

type Image struct {
	Caption []RichText `json:"caption"`
	Type    string     `json:"type"`
	File    File       `json:"file"`
}

type File struct {
	URL        string    `json:"url"`
	ExpiryTime time.Time `json:"expiry_time"`
}

type Embed struct {
	URL      string `json:"url"`
	Caption  string `json:"caption"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FullPage bool   `json:"full_page"`
}

type Todo struct {
	Text     []RichText `json:"text"`
	Checked  bool       `json:"checked"`
	Children []Block    `json:"children"`
}

type Toggle struct {
	Text     []RichText `json:"text"`
	Children []Block    `json:"children"`
}

type PageReference struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Bookmark struct {
	Caption []string `json:"caption"`
	URL     string   `json:"url"`
}
