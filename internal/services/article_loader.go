package services

import (
	"bytes"
	"path/filepath"
	"sort"
	"strings"

	embeded "github.com/Gsvd/gsvd.dev"
	"github.com/Gsvd/gsvd.dev/internal/models"
	"github.com/gernest/front"
	"github.com/mitchellh/mapstructure"
)

func LoadArticles() ([]models.Metadata, error) {
	var metadatas []models.Metadata
	articles, err := embeded.ContentFiles.ReadDir("internal/content")
	if err != nil {
		panic(err)
	}
	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)
	for _, article := range articles {
		if article.IsDir() {
			continue
		}

		fileContent, err := embeded.ContentFiles.ReadFile("internal/content/" + article.Name())
		if err != nil {
			return nil, err
		}
		f, _, err := m.Parse(bytes.NewReader(fileContent))
		if err != nil {
			return nil, err
		}

		metadata := &models.Metadata{}
		metadata.Slug = strings.TrimSuffix(article.Name(), filepath.Ext(article.Name()))
		if err := mapstructure.Decode(f, metadata); err != nil {
			return nil, err
		}
		metadatas = append(metadatas, *metadata)
	}

	sort.Slice(metadatas, func(i, j int) bool {
		return metadatas[i].Id > metadatas[j].Id
	})

	return metadatas, nil
}
