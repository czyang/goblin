package main

import (
	"html/template"
	"time"
)

type Post struct {
	Content    template.HTML
	MetaData   PostMeta
	CreateDate time.Time
	ModifyDate time.Time
}

type PostMeta struct {
	Title      string   `json:title`
	Tags       []string `json:tags`
	CreateDate string   `json:createDate`
	ModifyDate string   `json:modifyDate`
	Categories []string `json:categories`
	Permanent  string   `json:permanent`
	Author     string   `json:author`
}

type Page struct {
	Content    template.HTML
	MetaData   PostMeta
	CreateDate time.Time
	ModifyDate time.Time
}

type ArchivePage struct {
	Title string
	Posts []Post
}

type Nav_link struct {
	Title string `json:title`
	Url   string `json:url`
}

type Config struct {
	Nav_links     []Nav_link `json:nav_links`
	Twitter_Url   string     `json:twitter_url`
	Instagram_Url string     `json:instagram_url`
	Github_Url    string     `json:github_url`
	Weibo_Url     string     `json:weibo_url`
	Host  		  string	 `json:host`
	Port		  int		 `json:port`
}


type Posts []Post

type Pages []Page
