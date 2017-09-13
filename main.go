package main

import (
	"fmt"
	"net/http"
	"sort"
	"flag"
	"os"
	"strconv"
)

var workingPath string
var config Config

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Args error. Usage: goblin [working directory]")
		os.Exit(1)
	}
	workingPath = args[0]

	config = GetConfig()
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
	SpawnStaticPosts(posts)
	SpawnArchive(posts)
	SpawnStaticPages(pages)
	SpawnIndexPage()
	GenSiteMap(pages, posts)
	fmt.Println("Blog Spawn Success!")
	fmt.Println("http server started on [::]:" + strconv.Itoa(config.Port))

	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(":" + strconv.Itoa(config.Port), nil)
	checkError(err)
}
