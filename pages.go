package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func SpawnArchive(posts []Post) {
	pathString := path.Join(workingPath, "./static/pages/"+"archive"+".html")
	f, err := os.Create(pathString)
	checkError(err)
	t, err := template.ParseFiles("./tmpl/archive_layout.html")
	checkError(err)

	archivePage := ArchivePage{
		Title: "Archive",
		Posts: posts,
	}

	t.Execute(f, struct {
		Archive *ArchivePage
		Config  *Config
	}{&archivePage, &config})

	f.Close()
}

func GetPages() map[string]Page {
	pageMap := make(map[string]Page)
	files, _ := filepath.Glob("./source/pages/*")
	for _, f := range files {
		fileRead, _ := ioutil.ReadFile(f)
		jsonHeadBytes, mdBytes := ExtractMetaJSONStr(fileRead)
		postMeta := new(PostMeta)
		json.Unmarshal(jsonHeadBytes, &postMeta)
		articleContent := template.HTML(string(MarkdownBlog(mdBytes)))
		pageMap[postMeta.Permanent] = Page{Content: articleContent, MetaData: *postMeta,
			CreateDate: GetTime(postMeta.CreateDate), ModifyDate: GetTime(postMeta.ModifyDate)}
	}
	return pageMap
}

func SpawnStaticPages(pages []Page) {
	for _, v := range pages {
		pathString := path.Join(workingPath, "./static/pages/" + v.MetaData.Permanent + ".html")
		f, err := os.Create(pathString)
		checkError(err)
		t, err := template.ParseFiles("./tmpl/page_layout.html")
		checkError(err)

		t.Execute(f, struct {
			Page   *Page
			Config *Config
		}{&v, &config})

		f.Close()
	}
}

func SpawnIndexPage() {
	CopyFile(workingPath + "/static/pages/archive.html", workingPath + "/static/index.html")
}
