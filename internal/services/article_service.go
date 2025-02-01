package services

import (
	"bytes"
	"path/filepath"
	"sort"
	"strings"

	embeded "github.com/Gsvd/gsvd.dev"
	"github.com/Gsvd/gsvd.dev/internal/models"
	"github.com/Gsvd/gsvd.dev/internal/store"
	"github.com/gernest/front"
	"github.com/mitchellh/mapstructure"
)

func LoadMetadatas() ([]models.Metadata, error) {
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

func LoadComments(articleId int) ([]models.Comment, error) {
	var comments []models.Comment

	store := store.Get()
	rows, err := store.Query(`
		SELECT
			id,
			username,
			comment,
			approved,
			created_at
		FROM
			comments 
		WHERE
			post_id = ? 
			AND approved = TRUE 
		ORDER BY
			created_at DESC;
	`, articleId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.Id, &comment.Username, &comment.Comment, &comment.Approved, &comment.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func SaveComment(comment *models.Comment) error {
	store := store.Get()
	_, err := store.Exec(`
		INSERT INTO comments (post_id, username, comment, approved)
		VALUES (?, ?, ?, FALSE);
	`, comment.PostId, comment.Username, comment.Comment)

	return err
}
