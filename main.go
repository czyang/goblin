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

	config = GetConfig(inputPath)

	fmt.Println(0)
	postMap := GetPosts(inputPath)
	fmt.Println(1)
	var posts []Post
	for _, v := range postMap {
		posts = append(posts, v)
	}

	fmt.Println(1)
	sort.Sort(Posts(posts))
	CleanFolder(outputPath)
	fmt.Println(2)
	CreateFolder(outputPath+"/posts", 0755)
	CopyAssetsToStaticFolder(inputPath, outputPath)
	fmt.Println(3)
	SpawnStaticPosts(inputPath, outputPath, posts)
	SpawnIndex(outputPath, inputPath, posts)

	GenSiteMap(posts, outputPath)

	fmt.Println("Blog Spawn Success!")
}
