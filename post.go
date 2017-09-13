package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

)

// Sort posts.
func (s Posts) Len() int {
	return len(s)
}

func (s Posts) Less(i, j int) bool {
	return s[i].CreateDate.After(s[j].CreateDate)
}

func (s Posts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func ExtractMetaJSONStr(bytes []byte) ([]byte, []byte) {
	bracketCount := 0

	for index, byte := range bytes {
		if byte == '{' {
			bracketCount++
		} else if byte == '}' {
			bracketCount--
			if bracketCount == 0 {
				return bytes[:index+1], bytes[index+1:]
			}
		}
	}
	return nil, nil
}

func TrimTitleLine(b []byte) []byte {
	encounterHeader := false
	for i, v := range b {
		if v == '#' {
			encounterHeader = true
		}
		if encounterHeader && v == '\n' {
			return b[i+1:]
		}
	}
	return nil
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetTime(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	checkError(err)
	return t
}

func GetPosts() map[string]Post {
	postMap := make(map[string]Post)
	files, _ := filepath.Glob("./source/posts/*")
	for _, f := range files {
		fileRead, _ := ioutil.ReadFile(f)
		jsonHeadBytes, mdBytes := ExtractMetaJSONStr(fileRead)
		mdBytes = TrimTitleLine(mdBytes)
		postMeta := new(PostMeta)
		json.Unmarshal(jsonHeadBytes, &postMeta)
		articleContent := template.HTML(string(MarkdownBlog(mdBytes)))
		postMap[postMeta.Permanent] = Post{Content: articleContent, MetaData: *postMeta,
			CreateDate: GetTime(postMeta.CreateDate), ModifyDate: GetTime(postMeta.ModifyDate)}
	}
	return postMap
}

func SpawnStaticPosts(posts []Post) {
	for _, v := range posts {
		pathString := path.Join(workingPath, "./static/posts/" + v.MetaData.Permanent + ".html")
		f, err := os.Create(pathString)
		checkError(err)
		t, err := template.ParseFiles("./tmpl/post_layout.html")
		checkError(err)
		t.Execute(f, struct {
			Post   *Post
			Config *Config
		}{&v, &config})

		f.Close()
	}
}
