package main

import (
	"flag"
	"fmt"
	"sort"
)

var config Config

func main() {
	postsPath := flag.String("posts", "", "file path(string)")
	templatePath := flag.String("template", "", "file path(string)")
	outputPath := flag.String("output", "", "file path(string)")
	configPath := flag.String("config", "", "file path(string)")
	flag.Parse()

	config = GetConfig(*configPath)

	postMap := GetPosts(*postsPath)
	posts := make([]Post, len(postMap))
	i := 0
	for _, post := range postMap {
		posts[i] = post
		i++
	}

	sort.Sort(Posts(posts))
	CleanFolder(*outputPath)
	CreateFolder(*outputPath+"/posts", 0755)
	CopyAssetsAndPostMDToStaticFolder(*templatePath, *postsPath, *outputPath)
	SpawnStaticPosts(*templatePath, *outputPath, posts)
	SpawnIndex(*outputPath, *templatePath, posts)

	GenSiteMap(posts, *outputPath)

	fmt.Println("Blog Spawn Success!")
}
