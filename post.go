package main

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

// Posts sorting methods.
func (s Posts) Len() int           { return len(s) }
func (s Posts) Less(i, j int) bool { return s[i].CreateDate.After(s[j].CreateDate) }
func (s Posts) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Extracts the YAML header and content from input markdown file.
func ExtractMetaHeaderStr(input []byte) ([]byte, []byte) {
	parts := bytes.SplitN(input, []byte("---"), 3)
	if len(parts) < 3 {
		return nil, nil
	}
	return bytes.TrimSpace(parts[1]), bytes.TrimSpace(parts[2])
}

// Trim the markdown title from the content.
func TrimTitleLine(b []byte) []byte {
	for i, v := range b {
		if v == '#' {
			for v != '\n' {
				i++
				v = b[i]
			}
			return b[i+1:]
		}
	}
	return nil
}

// Parse date string into time.
func ParseDate(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	logFatalIfError(err)
	return t
}

// Fetches markdown posts from given path.
func GetPosts(postsPath string) map[string]Post {
	postMap := make(map[string]Post)
	files, _ := filepath.Glob(filepath.Join(postsPath, "*"))

	for _, f := range files {
		if filepath.Ext(f) != ".md" {
			continue
		}
		articleContent, postMeta := MarkdownFileToHTML(f)
		postMap[postMeta.Permanent] = Post{
			Content:    articleContent,
			MetaData:   *postMeta,
			CreateDate: ParseDate(postMeta.CreateDate),
			ModifyDate: ParseDate(postMeta.ModifyDate),
		}
	}
	return postMap
}

// Generates HTML static posts from given markdown posts.
func SpawnStaticPosts(templatePath string, outputPath string, posts []Post) {
	for _, post := range posts {
		f, err := os.Create(path.Join(outputPath, "posts", post.MetaData.Permanent+".html"))
		logFatalIfError(err)

		t, err := template.ParseFiles(filepath.Join(templatePath, "post_layout.html"))
		logFatalIfError(err)

		err = t.Execute(f, struct {
			Post   *Post
			Config *Config
		}{&post, &config})

		logFatalIfError(err)
		f.Close()
	}
}