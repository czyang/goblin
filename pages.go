package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// SpawnIndex generate archive page from template.
func SpawnIndex(outputPath string, templatePath string, posts []Post) {
	pathString := path.Join(outputPath, "/posts/"+"index"+".html")
	f, err := os.Create(pathString)
	checkError(err)
	t, err := template.ParseFiles(templatePath + "/archive_layout.html")
	checkError(err)

	archivePage := ArchivePage{
		Title: "Index",
		Posts: posts,
	}

	if err := t.Execute(f, struct {
		Archive *ArchivePage
		Config  *Config
	}{&archivePage, &config}); err != nil {
		panic(err)
	}

	f.Close()
}

// GetPages get pages.
func GetPages(inputPath string) map[string]Page {
	pageMap := make(map[string]Page)
	files, _ := filepath.Glob(inputPath + "/pages/*")
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
