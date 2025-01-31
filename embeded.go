package embeded

import "embed"

//go:embed public/*
var PublicFiles embed.FS

//go:embed dist/*
var DistFiles embed.FS

//go:embed internal/templates/*
var TemplateFiles embed.FS

//go:embed internal/content/*
var ArticleFiles embed.FS

//go:embed sitemap.xml
var SiteMapFile embed.FS
