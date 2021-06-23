package sitemapbuilder

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
	"webparser/parserapp/parser"
)

//SiteMap is the sitemap builder for the website
func SiteMap(urlStr string, depth int, l *log.Logger) []string {

	l.Println("**** Initiate site map build*****")
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= depth; i++ {
		q, nq = nq, make(map[string]struct{})
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			dup := url[:len(url)-1]
			if _, ok := seen[dup]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range getLinks(url) {
				nq[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	l.Println("**** finalizing site map build****")
	return ret
}

//getLinks will return thr filtered links
func getLinks(url string) []string {

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching URL response: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("Error parsing the html : %v", err)
	}

	//get the link nodes list
	linkNodes := parser.FilterLinkNodes(doc)
	reqUrl := resp.Request.URL

	baseUrl := &url2.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	//get formatted and filtered links
	return hrefs(linkNodes, base)
}

//hrefs will return the formatted links
func hrefs(linkNodes []*html.Node, base string) []string {

	//format the correct links
	var hrefs []string
	for _, l := range linkNodes {
		var href string
		for _, att := range l.Attr {
			if att.Key == "href" {
				href = att.Val
			}
		}
		//log.Printf("href : %v",href)
		switch {
		case strings.HasPrefix(href, "/"):
			hrefs = append(hrefs, base+href)
		case strings.HasPrefix(href, "http"):
			hrefs = append(hrefs, href)
		}
	}

	//filter the required links
	var filterLinks []string
	for _, link := range hrefs {
		if strings.HasPrefix(link, base) {
			filterLinks = append(filterLinks, link)
		}
	}
	return filterLinks
}
