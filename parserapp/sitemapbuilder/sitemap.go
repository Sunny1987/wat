package sitemapbuilder

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
	"sync"
	"webparser/parserapp/parser"
)

var wg sync.WaitGroup

//var wg1 sync.WaitGroup

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
			l.Printf("Adding to seen: %v", url)
			seen[url] = struct{}{}
			for _, link := range getLinks(url, l) {
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
func getLinks(url string, l *log.Logger) []string {
	l.Println("*****Starting to get filtered links****")
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching URL response: %v", err)
	}
	defer resp.Body.Close()

	l.Println("****parsing the document*****")
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("Error parsing the html : %v", err)
	}

	l.Println("****getting the linkNodes*****")

	//get the link nodes list
	linkNodes := parser.FilterLinkNodes(doc)
	reqUrl := resp.Request.URL

	l.Println("*** build base url****")
	baseUrl := &url2.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	//get formatted and filtered links
	return hrefs(linkNodes, base, l)
}

//hrefs will return the formatted links
func hrefs(linkNodes []*html.Node, base string, l *log.Logger) []string {

	l.Println("***Formatting the link nodes****")

	//format the correct links
	var hrefs []string
	wg.Add(len(linkNodes))
	for _, l := range linkNodes {

		go func(l *html.Node) {
			defer wg.Done()
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

		}(l)

	}

	wg.Wait()
	//filter the required links
	var filterLinks []string
	//wg.Add(len(hrefs))
	for _, link := range hrefs {
		//go func(link string) {
		//	defer wg.Done()
		if strings.HasPrefix(link, base) {
			filterLinks = append(filterLinks, link)
		}
		//}(link)

	}
	//wg.Wait()
	l.Println("***Completed formatting the link nodes****")
	return filterLinks
}
