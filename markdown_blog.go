package main

import "github.com/russross/blackfriday"

const blogHtmlFlags = 0 |
	blackfriday.HTML_USE_XHTML |
	blackfriday.HTML_USE_SMARTYPANTS |
	blackfriday.HTML_SMARTYPANTS_FRACTIONS |
	blackfriday.HTML_SMARTYPANTS_DASHES |
	blackfriday.HTML_SMARTYPANTS_LATEX_DASHES |
	blackfriday.HTML_SMARTYPANTS_ANGLED_QUOTES |
	blackfriday.HTML_HREF_TARGET_BLANK

const commonExtensions = 0 |
	blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
	blackfriday.EXTENSION_TABLES |
	blackfriday.EXTENSION_FENCED_CODE |
	blackfriday.EXTENSION_AUTOLINK |
	blackfriday.EXTENSION_STRIKETHROUGH |
	blackfriday.EXTENSION_SPACE_HEADERS |
	blackfriday.EXTENSION_HEADER_IDS |
	blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
	blackfriday.EXTENSION_DEFINITION_LISTS

// Custom Markdown process function for blog, inherited from MarkdownCommon.
func MarkdownBlog(input []byte) []byte {
	// set up the HTML renderer
	renderer := blackfriday.HtmlRenderer(blogHtmlFlags, "", "")
	return blackfriday.MarkdownOptions(input, renderer, blackfriday.Options{
		Extensions: commonExtensions})
}
