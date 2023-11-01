package embeded

import "embed"

//go:embed public/*
var PublicFiles embed.FS

//go:embed dist/*
var DistFiles embed.FS

//go:embed templates/*
var TemplateFiles embed.FS

//go:embed articles/*
var ArticleFiles embed.FS
