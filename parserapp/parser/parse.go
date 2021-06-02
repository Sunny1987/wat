package parser

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"sync"
	"webparser/analyzerapp"
)

var wg sync.WaitGroup

type Link struct {
	Href string
	Text string
}


//Parse will parse the response body
func (ml *MyLogger) Parse(r io.Reader) analyzerapp.Response {
	//io.Copy(os.Stdout,r)
	doc, err := html.Parse(r)
	if err != nil {
		ml.l.Printf("Error parsing the html : %v", err)
	}

	//get all nodes in map
	nodeMap := collectNodes(doc, ml.l)

	//initiate analysis
	ml.l.Println("Initiating WCAG 2.1 analysis.....")
	results := analyzerapp.Analysis(ml.l).ApplyRules(nodeMap)
	return results

}

//collectNodes collects all the nodes and stores in nodeMap
func collectNodes(doc *html.Node, l *log.Logger) map[string][]*html.Node {
	//List of link nodes
	nodeMap := make(map[string][]*html.Node)
	wg.Add(1)
	go func() {
		l.Println("Collecting all links...")
		linkNodes := filterLinkNodes(doc)
		if len(linkNodes) > 0 {
			nodeMap["linkNodes"] = linkNodes
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		//List of Divs
		l.Println("Collecting all divs...")
		divNodes := filterDivNodes(doc)
		if len(divNodes) > 0 {
			nodeMap["divNodes"] = divNodes
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of Paragraphs
		l.Println("Collecting all paragraphs...")
		paraNodes := filterParaNodes(doc)
		if len(paraNodes) > 0 {
			nodeMap["paraNodes"] = paraNodes
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of spans
		l.Println("Collecting all spans...")
		spanNodes := filterSpanNodes(doc)
		if len(spanNodes) > 0 {
			nodeMap["spanNodes"] = spanNodes
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of h1
		l.Println("Collecting all h1...")
		h1Nodes := filterH1Nodes(doc)
		if len(h1Nodes) > 0 {
			nodeMap["h1Nodes"] = h1Nodes
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of h2
		l.Println("Collecting all h2...")
		h2Nodes := filterH2Nodes(doc)
		if len(h2Nodes) > 0 {
			nodeMap["h2Nodes"] = h2Nodes
		}

		wg.Done()

	}()

	wg.Add(1)
	go func() {
		//List of h3
		l.Println("Collecting all h3...")
		h3Nodes := filterH3Nodes(doc)
		if len(h3Nodes) > 0 {
			nodeMap["h3Nodes"] = h3Nodes
		}

		wg.Done()

	}()

	wg.Add(1)
	go func() {
		//List of h4
		l.Println("Collecting all h4...")
		h4Nodes := filterH4Nodes(doc)
		nodeMap["h4Nodes"] = h4Nodes
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		//List of h5
		l.Println("Collecting all h5...")
		h5Nodes := filterH5Nodes(doc)
		nodeMap["h5Nodes"] = h5Nodes
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of h6
		l.Println("Collecting all h6...")
		h6Nodes := filterH6Nodes(doc)
		nodeMap["h6Nodes"] = h6Nodes
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of image
		l.Println("Collecting all images...")
		imgNodes := filterImageNodes(doc)
		nodeMap["imgNodes"] = imgNodes
		wg.Done()
	}()

	wg.Wait()
	return nodeMap

}
