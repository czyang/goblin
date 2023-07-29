package main

import (
	"html/template"
	"time"
)

// Post post struct
type Post struct {
	Content    template.HTML
	MetaData   PostMeta
	CreateDate time.Time
	ModifyDate time.Time
}

// PostMeta meta data struct of the Post
type PostMeta struct {
	Title      string   `yaml:"title"`
	Tags       []string `yaml:"tags"`
	CreateDate string   `yaml:"createDate"`
	ModifyDate string   `yaml:"modifyDate"`
	Categories []string `yaml:"categories"`
	Permanent  string   `yaml:"permanent"`
	Author     string   `yaml:"author"`
}

// Page page struct.
type Page struct {
	Content    template.HTML
	MetaData   PostMeta
	CreateDate time.Time
	ModifyDate time.Time
}

// ArchivePage archive page struct.
type ArchivePage struct {
	Title string
	Posts []Post
}

// NavLink link struct on the page header.
type NavLink struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Config config struct
type Config struct {
	NavLinks     []NavLink `json:"nav_links"`
	TwitterURL   string    `json:"twitter_url"`
	InstagramURL string    `json:"instagram_url"`
	GithubURL    string    `json:"github_url"`
	WeiboURL     string    `json:"weibo_url"`
	Host         string    `json:"host"`
	Port         int       `json:"port"`
}

// Posts post map
type Posts []Post

// Pages page map
type Pages []Page
