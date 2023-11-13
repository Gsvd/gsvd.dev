package internal

import (
	"bytes"
	"path/filepath"
	"sort"
	"strings"

	embeded "github.com/Gsvd/gsvd.dev"
	"github.com/gernest/front"
	"github.com/mitchellh/mapstructure"
)

func LoadArticlesMetadata() ([]ArticleMetadata, error) {
	var articlesMetadata []ArticleMetadata
	articles, err := embeded.ArticleFiles.ReadDir("articles")
	if err != nil {
		panic(err)
	}
	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)
	for _, article := range articles {
		if article.IsDir() {
			continue
		}

		fileContent, err := embeded.ArticleFiles.ReadFile("articles/" + article.Name())
		if err != nil {
			return nil, err
		}
		f, _, err := m.Parse(bytes.NewReader(fileContent))
		if err != nil {
			return nil, err
		}

		metadata := &ArticleMetadata{}
		metadata.Slug = strings.TrimSuffix(article.Name(), filepath.Ext(article.Name()))
		if err := mapstructure.Decode(f, metadata); err != nil {
			return nil, err
		}
		articlesMetadata = append(articlesMetadata, *metadata)
	}

	sort.Slice(articlesMetadata, func(i, j int) bool {
		return articlesMetadata[i].Id > articlesMetadata[j].Id
	})

	return articlesMetadata, nil
}
