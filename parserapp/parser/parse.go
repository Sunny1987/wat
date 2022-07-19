package parser

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"sync"
	"webparser/analyzerapp"
)

var (
	wg      sync.WaitGroup
	mu      sync.RWMutex
	cssList []string
)

//Parse will parse the response body
func (ml *MyLogger) Parse(r io.Reader, base string) analyzerapp.Response {
	//io.Copy(os.Stdout,r)
	doc, err := html.Parse(r)
	if err != nil {
		ml.l.Printf("Error parsing the html : %v", err)
	}

	//get all nodes in map
	nodeMap := collectNodes(doc, ml.l, base)

	//initiate analysis
	ml.l.Println("Initiating WCAG 2.1 analysis.....")
	results := analyzerapp.Analysis(ml.l, ml.req, ml.person).ApplyRules(nodeMap, cssList)
	return results

}

//collectNodes collects all the nodes and stores in nodeMap
func collectNodes(doc *html.Node, l *log.Logger, base string) map[string][]*html.Node {
	//List of link nodes
	nodeMap := make(map[string][]*html.Node)

	wg.Add(24)
	go getNode(FilterAnchorNodes, "anchorNodes", doc, l, nodeMap)
	go getNode(filterDivNodes, "divNodes", doc, l, nodeMap)
	go getNode(filterParaNodes, "paraNodes", doc, l, nodeMap)
	go getNode(filterSpanNodes, "spanNodes", doc, l, nodeMap)
	go getNode(filterH1Nodes, "h1Nodes", doc, l, nodeMap)
	go getNode(filterH2Nodes, "h2Nodes", doc, l, nodeMap)
	go getNode(filterH3Nodes, "h3Nodes", doc, l, nodeMap)
	go getNode(filterH4Nodes, "h4Nodes", doc, l, nodeMap)
	go getNode(filterH5Nodes, "h5Nodes", doc, l, nodeMap)
	go getNode(filterH6Nodes, "h6Nodes", doc, l, nodeMap)
	go getNode(filterImageNodes, "imgNodes", doc, l, nodeMap)
	go getNode(filterInputNodes, "inputNodes", doc, l, nodeMap)
	go getNode(filterButtonNodes, "buttonNodes", doc, l, nodeMap)
	go getNode(filterVideoNodes, "videoNodes", doc, l, nodeMap)
	go getNode(filterAudioNodes, "audioNodes", doc, l, nodeMap)
	go getNode(filterSelectNodes, "selectNodes", doc, l, nodeMap)
	go getNode(filterTextAreaNodes, "textareaNodes", doc, l, nodeMap)
	go getNode(filterIframeNodes, "iframeNodes", doc, l, nodeMap)
	go getNode(filterAreaNodes, "areaNodes", doc, l, nodeMap)
	go getNode(filterObjectNodes, "objectNodes", doc, l, nodeMap)
	go getNode(filterEmbedNodes, "embedNodes", doc, l, nodeMap)
	go getNode(filterTrackNodes, "trackNodes", doc, l, nodeMap)
	go getNode(filterAppletNodes, "appletNodes", doc, l, nodeMap)
	go getNode(filterPreNodes, "preNodes", doc, l, nodeMap)

	wg.Add(1)
	go func(base string) {
		l.Println("Collecting all links...")
		defer wg.Done()
		//var cssLinks []string
		//cssLinks = nil
		if base != "" {
			//get all links
			linkNodes := FilterLinkNodes(doc)
			l.Printf("base link: %v", base)

			//collect CSS links
			cssLinks := HrefLinks(filterCSSLinks(linkNodes), base, l)
			for _, link := range cssLinks {
				l.Println(link)
				readCSSLinks(link, l)
			}

			if len(linkNodes) > 0 {
				mu.Lock()
				nodeMap["linkNodes"] = linkNodes
				mu.Unlock()
			}
		}

	}(base)

	wg.Wait()
	return nodeMap

}

//getNode function is a common function to retrieve the node
func getNode(fn func(node *html.Node) []*html.Node, nodeName string, doc *html.Node, l *log.Logger, nodeMap map[string][]*html.Node) {
	defer wg.Done()
	l.Printf("Collecting all %v...", nodeName)
	nodes := fn(doc)
	if len(nodes) > 0 {
		mu.Lock()
		nodeMap[nodeName] = nodes
		mu.Unlock()
	}
}
