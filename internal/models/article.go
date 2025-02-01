package models

import "html/template"

type Article struct {
	Metadata Metadata
	Content  template.HTML
	Comments []Comment
}
