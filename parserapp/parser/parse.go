package parser

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"sync"
	"webparser/analyzerapp"
)

var wg sync.WaitGroup
var mu sync.RWMutex

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
			mu.Lock()
			nodeMap["linkNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		//List of Divs
		l.Println("Collecting all divs...")
		divNodes := filterDivNodes(doc)
		if len(divNodes) > 0 {
			mu.Lock()
			nodeMap["divNodes"] = divNodes
			mu.Unlock()
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of Paragraphs
		l.Println("Collecting all paragraphs...")
		paraNodes := filterParaNodes(doc)
		if len(paraNodes) > 0 {
			mu.Lock()
			nodeMap["paraNodes"] = paraNodes
			mu.Unlock()
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of spans
		l.Println("Collecting all spans...")
		spanNodes := filterSpanNodes(doc)
		if len(spanNodes) > 0 {
			mu.Lock()
			nodeMap["spanNodes"] = spanNodes
			mu.Unlock()
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of h1
		l.Println("Collecting all h1...")
		h1Nodes := filterH1Nodes(doc)
		if len(h1Nodes) > 0 {
			mu.Lock()
			nodeMap["h1Nodes"] = h1Nodes
			mu.Unlock()
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of h2
		l.Println("Collecting all h2...")
		h2Nodes := filterH2Nodes(doc)
		if len(h2Nodes) > 0 {
			mu.Lock()
			nodeMap["h2Nodes"] = h2Nodes
			mu.Unlock()
		}

		wg.Done()

	}()

	wg.Add(1)
	go func() {
		//List of h3
		l.Println("Collecting all h3...")
		h3Nodes := filterH3Nodes(doc)
		if len(h3Nodes) > 0 {
			mu.Lock()
			nodeMap["h3Nodes"] = h3Nodes
			mu.Unlock()
		}

		wg.Done()

	}()

	wg.Add(1)
	go func() {
		//List of h4
		l.Println("Collecting all h4...")
		h4Nodes := filterH4Nodes(doc)

		if len(h4Nodes) > 0 {
			mu.Lock()
			nodeMap["h4Nodes"] = h4Nodes
			mu.Unlock()
		}

		wg.Done()

	}()

	wg.Add(1)
	go func() {
		//List of h5
		l.Println("Collecting all h5...")
		h5Nodes := filterH5Nodes(doc)
		if len(h5Nodes) > 0 {
			mu.Lock()
			nodeMap["h5Nodes"] = h5Nodes
			mu.Unlock()
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of h6
		l.Println("Collecting all h6...")
		h6Nodes := filterH6Nodes(doc)
		if len(h6Nodes) > 0 {
			mu.Lock()
			nodeMap["h6Nodes"] = h6Nodes
			mu.Unlock()
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//List of image
		l.Println("Collecting all images...")
		imgNodes := filterImageNodes(doc)
		if len(imgNodes) > 0 {
			mu.Lock()
			nodeMap["imgNodes"] = imgNodes
			mu.Unlock()
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all inputs...")
		linkNodes := filterInputNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["inputNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all buttons...")
		linkNodes := filterButtonNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["buttonNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all videos...")
		linkNodes := filterVideoNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["videoNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all audio...")
		linkNodes := filterAudioNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["audioNodes"] = linkNodes
			mu.Unlock()

		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all selects...")
		linkNodes := filterSelectNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["selectNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all textareas...")
		linkNodes := filterTextAreaNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["textareaNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all iframes...")
		linkNodes := filterIframeNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["iframeNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all areas...")
		linkNodes := filterAreaNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["areaNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all objects...")
		linkNodes := filterObjectNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["objectNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all embeds...")
		linkNodes := filterEmbedNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["embedNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all tracks...")
		linkNodes := filterTrackNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["trackNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Add(1)
	go func() {
		l.Println("Collecting all applets...")
		linkNodes := filterAppletNodes(doc)
		if len(linkNodes) > 0 {
			mu.Lock()
			nodeMap["appletNodes"] = linkNodes
			mu.Unlock()
		}
		wg.Done()

	}()

	wg.Wait()
	return nodeMap

}
