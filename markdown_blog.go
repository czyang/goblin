package main

import (
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

// MarkdownFileToHTML Generate HTML string from giving file
func MarkdownFileToHTML(fileName string) (template.HTML, *PostMeta) {
	fileRead, _ := ioutil.ReadFile(fileName)
	headerBytes, mdBytes := ExtractMetaHeaderStr(fileRead)

	// mdBytes = TrimTitleLine(mdBytes)
	postMeta := new(PostMeta)
	err := yaml.Unmarshal(headerBytes, &postMeta)
	if err != nil {
		log.Println(fileName)
		panic(err)
	}

	return template.HTML(blackfriday.Run(mdBytes)), postMeta
}
