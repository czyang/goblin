package main

import (
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
)

func GenSiteMap(pages Pages, posts Posts) {
	items := make([]SiteMapItem, 0)

	for _, v := range pages {
		item := SiteMapItem{Loc:"/pages/" + v.MetaData.Permanent + ".html", Changefreq: "weekly", Mobile: true}
		items = append(items, item)
	}

	for _, v := range posts {
		item := SiteMapItem{Loc:"/posts/" + v.MetaData.Permanent + ".html", Changefreq: "weekly", Mobile: true}
		items = append(items, item)
	}

	//index.html
	item := SiteMapItem{Loc:"/", Changefreq: "weekly", Mobile: true}
	items = append(items, item)

	sm := stm.NewSitemap()
	sm.SetDefaultHost(config.Host)
	sm.SetSitemapsPath("/static")
	sm.SetPublicPath("./")

	sm.Create()

	for _, v := range items {
		sm.Add(stm.URL{"loc": v.Loc, "changefreq": v.Changefreq, "mobile": v.Mobile})
	}

	sm.Finalize().PingSearchEngines()
}