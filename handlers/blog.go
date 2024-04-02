package handlers

import (
	"bytes"
	"database/sql"
	"fmt"
	embeded "github.com/Gsvd/gsvd.dev"
	"github.com/Gsvd/gsvd.dev/internal"
	"github.com/gernest/front"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Comment struct {
	Id               int
	Author           string
	Content          string
	Created          time.Time
	CreatedFormatted string
}

func BlogHandler(c *fiber.Ctx) error {
	articlesMetadata, err := internal.LoadArticlesMetadata()
	if err != nil {
		panic(err)
	}
	return c.Render("templates/blog", fiber.Map{
		"Title":     "Blog Articles - Gsvd",
		"Articles":  articlesMetadata,
		"Canonical": "blog",
	}, "templates/layouts/main")
}

func BlogPostHandler(c *fiber.Ctx) error {
	filename := fmt.Sprintf("articles/%s.md", c.Params("title"))

	fileContent, err := embeded.ArticleFiles.ReadFile(filename)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)
	f, body, err := m.Parse(bytes.NewReader(fileContent))
	if err != nil {
		panic(err)
	}

	metadata := &internal.ArticleMetadata{}
	if err := mapstructure.Decode(f, metadata); err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	comments := make([]Comment, 0)

	results, err := db.Query("SELECT id, author, content, created_at FROM comments WHERE article_id = ? AND approved = 1;", metadata.Id)
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()

	for results.Next() {
		var c Comment
		err := results.Scan(&c.Id, &c.Author, &c.Content, &c.Created)
		if err != nil {
			panic(err.Error())
		}
		c.CreatedFormatted = c.Created.Format("2006-01-02 15:04:05")
		comments = append(comments, c)
	}

	htmlContent := blackfriday.Run([]byte(body))

	article := internal.Article{
		Metadata: *metadata,
		Content:  template.HTML(htmlContent),
	}

	return c.Render("templates/post", fiber.Map{
		"Title":     metadata.Title + " - Gsvd",
		"Article":   article,
		"Canonical": "blog/" + metadata.Slug,
		"Comments":  comments,
	}, "templates/layouts/post")
}

func BlogPostCommentHandler(c *fiber.Ctx) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	articleId := c.Params("id")
	articleSlug := c.Params("title")
	author := c.FormValue("author")
	content := c.FormValue("content")

	if author == "" || content == "" || articleId == "" || articleSlug == "" {
		return c.Redirect(fmt.Sprintf("/blog/%s", articleSlug))
	}

	_, err = db.Exec("INSERT INTO comments (article_id, author, content) VALUES (?, ?, ?);", articleId, author, content)
	if err != nil {
		log.Println(err.Error())
	}

	return c.Redirect(fmt.Sprintf("/blog/%s", articleSlug))
}
