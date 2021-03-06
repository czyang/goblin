package main

import (
	"flag"
	"fmt"
	"sort"
)

var config Config

func main() {
	postsPathPtr := flag.String("posts", "", "file path(string)")
	templatePathPtr := flag.String("template", "", "file path(string)")
	outputPathPtr := flag.String("output", "", "file path(string)")
	flag.Parse()
	postsPath := *postsPathPtr
	templatePath := *templatePathPtr
	outputPath := *outputPathPtr

	config = GetConfig()

	postMap := GetPosts(postsPath)
	var posts []Post
	for _, v := range postMap {
		posts = append(posts, v)
	}

	sort.Sort(Posts(posts))
	CleanFolder(outputPath)
	CreateFolder(outputPath+"/posts", 0755)
	CopyAssetsToStaticFolder(templatePath, postsPath, outputPath)
	SpawnStaticPosts(templatePath, outputPath, posts)
	SpawnIndex(outputPath, templatePath, posts)

	GenSiteMap(posts, outputPath)

	fmt.Println("Blog Spawn Success!")
}
