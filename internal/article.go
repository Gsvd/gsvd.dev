package internal

import "html/template"

type Article struct {
	Metadata ArticleMetadata
	Content  template.HTML
}
