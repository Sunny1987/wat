package parserhandler

import (
	"net/http"
	"sync"
	"time"
	"webparser/analyzerapp"
	"webparser/parserapp/sitemapbuilder"
)

var wg sync.WaitGroup

//GetURLResp will return the WCAG2.1 guidelines result for a URL
func (n *NewLogger) GetURLResp(rw http.ResponseWriter, r *http.Request) {

	//track execution time for scan
	timeStart := time.Now()

	//get the request from middleware
	req := r.Context().Value(KeyUser{}).(*MyURLReq)

	//get the list of links from sitemap
	links, base := sitemapbuilder.SiteMap(req.URLFromReq, req.MaxDepth, n.l)
	n.l.Println("***** site map completed*****")
	var finalResult []analyzerapp.Response

	//scan based on depth
	wg.Add(len(links))
	for i, link := range links {

		go func(link string) {
			defer wg.Done()

			//setup req object
			reqMod := &MyURLReq{}
			reqMod.URLFromReq = link
			reqMod.MaxDepth = req.MaxDepth

			n.l.Printf("Link# %v : %v ", i, reqMod.URLFromReq)
			//start scan for url
			results := startScan(reqMod, n.l, base)
			//mu.Lock()
			finalResult = append(finalResult, results)
			//mu.Unlock()

		}(link)

	}

	wg.Wait()
	//print the response
	PrintResponse(finalResult, rw, n.l)

	n.l.Printf("Query completed in %v\n", time.Since(timeStart))

}
