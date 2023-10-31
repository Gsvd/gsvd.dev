package internal

import (
	"bytes"
	"github.com/gernest/front"
	"github.com/mitchellh/mapstructure"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func LoadArticlesMetadata() ([]ArticleMetadata, error) {
	var articlesMetadata []ArticleMetadata
	articles, err := os.ReadDir("./articles/")
	if err != nil {
		panic(err)
	}
	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)
	for _, article := range articles {
		if article.IsDir() {
			continue
		}

		fileContent, err := os.ReadFile("./articles/" + article.Name())
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
		return articlesMetadata[i].Order > articlesMetadata[j].Order
	})

	return articlesMetadata, nil
}
