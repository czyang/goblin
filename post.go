package main

import (
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
	"bytes"
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

// ExtractMetaHeaderStr extract the page's meta data from the yaml string which
// on the header of the article.
func ExtractMetaHeaderStr(input []byte) ([]byte, []byte) {
    separator := []byte("---")
    parts := bytes.SplitN(input, separator, 3)

    if len(parts) < 3 {
        // Not enough separators were found.
        return nil, nil
    }

    // Trim spaces before returning
    return bytes.TrimSpace(parts[1]), bytes.TrimSpace(parts[2])
}

// TrimTitleLine trim off the meta data string.
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

// GetTime get the time from string.
func GetTime(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	checkError(err)
	return t
}

// GetPosts get the posts from input/posts. All posts are markdown format.
func GetPosts(postsPath string) map[string]Post {
	postMap := make(map[string]Post)
	files, _ := filepath.Glob(postsPath + "/*")

	for _, f := range files {
		articleContent, postMeta := MarkdownFileToHTML(f)
		extension := filepath.Ext(f)
		if extension != ".md" {
			continue
		}
		postMap[postMeta.Permanent] = Post{Content: articleContent, MetaData: *postMeta,
			CreateDate: GetTime(postMeta.CreateDate), ModifyDate: GetTime(postMeta.ModifyDate)}
	}
	return postMap
}

// SpawnStaticPosts generate all static posts which are HTML.
func SpawnStaticPosts(templatePath string, outputPath string, posts []Post) {
	for _, v := range posts {
		pathString := path.Join(outputPath, "/posts/"+v.MetaData.Permanent+".html")
		f, err := os.Create(pathString)
		checkError(err)
		t, err := template.ParseFiles(templatePath + "/post_layout.html")
		checkError(err)
		if err := t.Execute(f, struct {
			Post   *Post
			Config *Config
		}{&v, &config}); err != nil {
			panic(err)
		}
		f.Close()
	}
}
