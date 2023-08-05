package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/snabb/sitemap"
)

// GenSiteMap generate sitemap.
func GenSiteMap(posts Posts, outputPath string) {
	sm := sitemap.New()

	t := time.Now().UTC()
	for _, v := range posts {
		sm.Add(&sitemap.URL{
			Loc:        config.Host + "/posts/" + v.MetaData.Permanent + ".html",
			LastMod:    &t,
			ChangeFreq: sitemap.Weekly,
		})
	}

	//index.html
	sm.Add(&sitemap.URL{
		Loc:        config.Host + "/",
		LastMod:    &t,
		ChangeFreq: sitemap.Weekly,
	})

	pathString := path.Join(outputPath, "/posts/sitemap.xml")
	f, err := os.Create(pathString)
	logFatalIfError(err)
	sm.WriteTo(f)

	PingSearchEngines()
}

// PingSearchEngines requests some ping server from it calls Sitemap.PingSearchEngines.
// Modify from: github.com/ikeikeikeike/go-sitemap-generator/stm"
func PingSearchEngines(urls ...string) {
	urls = append(urls, []string{
		"http://www.google.com/webmasters/tools/ping?sitemap=%s",
		"http://www.bing.com/webmaster/ping.aspx?siteMap=%s",
	}...)

	bufs := len(urls)
	does := make(chan string, bufs)
	client := http.Client{Timeout: time.Duration(5 * time.Second)}

	for _, url := range urls {
		go func(baseurl string) {
			pingURL := fmt.Sprintf(baseurl, config.Host+"/sitemap.xml")
			println("Ping now:", pingURL)

			resp, err := client.Get(pingURL)
			if err != nil {
				does <- fmt.Sprintf("[E] Ping failed: %s (URL:%s)",
					err, pingURL)
				return
			}
			defer resp.Body.Close()

			does <- fmt.Sprintf("Successful ping of `%s`", pingURL)
		}(url)
	}

	for i := 0; i < bufs; i++ {
		println(<-does)
	}
}
