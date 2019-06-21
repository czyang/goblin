package main

import (
	"encoding/json"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
)

// MarkdownFileToHTML Generate HTML string from giving file
func MarkdownFileToHTML(fileName string) (template.HTML, *PostMeta) {
	fileRead, _ := ioutil.ReadFile(fileName)
	jsonHeadBytes, mdBytes := ExtractMetaJSONStr(fileRead)
	mdBytes = TrimTitleLine(mdBytes)
	postMeta := new(PostMeta)
	if err := json.Unmarshal(jsonHeadBytes, &postMeta); err {
		checkError(err)
	}
	return template.HTML(blackfriday.Run(mdBytes)), postMeta
}
