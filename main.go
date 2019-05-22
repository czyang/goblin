package main

import (
	"flag"
	"fmt"
	"sort"
)

var config Config

func main() {
	inputPathPtr := flag.String("input", "", "file path(string)")
	outputPathPtr := flag.String("output", "", "file path(string)")
	flag.Parse()
	inputPath := *inputPathPtr
	outputPath := *outputPathPtr
	fmt.Println("inputPathPtr:", inputPath, "outputPathPtr:", outputPath)

	config = GetConfig()
	postMap := GetPosts(inputPath)
	var posts []Post
	for _, v := range postMap {
		posts = append(posts, v)
	}
	sort.Sort(Posts(posts))
	CleanFolder(outputPath)
	CreateFolder(outputPath+"/posts", 0755)
	CopyAssetsToStaticFolder(inputPath, outputPath)
	SpawnStaticPosts(inputPath, outputPath, posts)
	SpawnIndex(outputPath, inputPath, posts)

	GenSiteMap(posts, outputPath)

	fmt.Println("Blog Spawn Success!")
}
