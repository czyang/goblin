package main

import (
	"fmt"
	"net/http"
	"sort"
)

func main() {
	config := GetConfig()

	postMap := GetPosts()
	var posts []Post
	for _, v := range postMap {
		posts = append(posts, v)
	}
	sort.Sort(Posts(posts))

	pageMap := GetPages()
	var pages []Page
	for _, v := range pageMap {
		pages = append(pages, v)
	}
	CleanStaticFolder()
	CreateFolder("./static/posts", 0777)
	CreateFolder("./static/pages", 0777)
	CopyAssetsToStaticFolder()
	SpawnStaticPosts(posts, config)
	SpawnArchive(posts, config)
	SpawnStaticPages(pages, config)
	SpawnIndexPage()
	fmt.Println("Blog Spawn Success!")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(":8001", nil)
	checkError(err)
}
