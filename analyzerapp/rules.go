package analyzerapp

import (
	"golang.org/x/net/html"
	"log"
	"sync"
)

//MyAnalysisLog is the logger for analysisapp
type MyAnalysisLog struct {
	l      *log.Logger
	req    interface{}
	person string
}

//Analysis will return the MyAnalysisLog logger
func Analysis(l *log.Logger, req interface{}, person string) *MyAnalysisLog {
	return &MyAnalysisLog{l: l, req: req, person: person}
}

var wg sync.WaitGroup
var mu sync.RWMutex

//ApplyRules will initiate rule application
func (l *MyAnalysisLog) ApplyRules(nodeMap map[string][]*html.Node, cssList []string) Response {
	l.l.Println("Initiating Rules Check....")
	var ruleResults Response

	//Add the requestR
	ruleResults.Request = l.req

	//Add the credential Details
	ruleResults.Person = l.person

	//div analysis
	wg.Add(1)
	go func() {
		myMap := divAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["div"]
		mu.RUnlock()
		ruleResults.DivResults = myList
		wg.Done()
	}()

	//button analysis
	wg.Add(1)
	go func() {
		myMap := buttonAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["button"]
		mu.RUnlock()
		ruleResults.ButtonResults = myList
		wg.Done()
	}()

	//input analysis
	wg.Add(1)
	go func() {
		myMap := inputAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["input"]
		mu.RUnlock()
		ruleResults.InputResults = myList
		wg.Done()
	}()

	//images analysis
	wg.Add(1)
	go func() {
		myMap := imagesAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["image"]
		mu.RUnlock()
		ruleResults.ImageResults = myList
		wg.Done()
	}()

	//video analysis
	wg.Add(1)
	go func() {
		myMap := videoAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["video"]
		mu.RUnlock()
		ruleResults.VideoResults = myList
		wg.Done()
	}()

	//audio analysis
	wg.Add(1)
	go func() {
		myMap := audioAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["audio"]
		mu.RUnlock()
		ruleResults.AudioResults = myList
		wg.Done()
	}()

	//textarea analysis
	wg.Add(1)
	go func() {
		myMap := textareaAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["textarea"]
		mu.RUnlock()
		ruleResults.TextareaResults = myList
		wg.Done()
	}()

	//select analysis
	wg.Add(1)
	go func() {
		myMap := selectAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["select"]
		mu.RUnlock()
		ruleResults.SelectResults = myList
		wg.Done()
	}()

	//iframe analysis
	wg.Add(1)
	go func() {
		myMap := iframeAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["iframe"]
		mu.RUnlock()
		ruleResults.IframeResults = myList
		wg.Done()
	}()

	//anchor analysis
	wg.Add(1)
	go func() {
		myMap := linkAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["anchor"]
		mu.RUnlock()
		ruleResults.LinkResults = myList
		wg.Done()
	}()

	//area analysis
	wg.Add(1)
	go func() {
		myMap := areaAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["area"]
		mu.RUnlock()
		ruleResults.AreaResults = myList
		wg.Done()
	}()

	//object analysis
	wg.Add(1)
	go func() {
		myMap := objectAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["object"]
		mu.RUnlock()
		ruleResults.ObjectResults = myList
		wg.Done()
	}()

	//embed analysis
	wg.Add(1)
	go func() {
		myMap := embedAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["embed"]
		mu.RUnlock()
		ruleResults.EmbedResults = myList
		wg.Done()
	}()

	//track analysis
	wg.Add(1)
	go func() {
		myMap := trackAnalysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["track"]
		mu.RUnlock()
		ruleResults.TrackResults = myList
		wg.Done()
	}()

	//css analysis
	wg.Add(1)
	go func() {
		defer wg.Done()
		cssAnalysis(l.l, cssList)
	}()

	//h1 analysis
	wg.Add(1)
	go func() {
		defer wg.Done()
		myMap := h1Analysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["h1"]
		mu.RUnlock()
		ruleResults.H1Results = myList

	}()

	//h2 analysis
	wg.Add(1)
	go func() {
		defer wg.Done()
		myMap := h2Analysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["h2"]
		mu.RUnlock()
		ruleResults.H2Results = myList

	}()

	//h3 analysis
	wg.Add(1)
	go func() {
		defer wg.Done()
		myMap := h3Analysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["h3"]
		mu.RUnlock()
		ruleResults.H3Results = myList

	}()

	//h4 analysis
	wg.Add(1)
	go func() {
		defer wg.Done()
		myMap := h4Analysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["h4"]
		mu.RUnlock()
		ruleResults.H4Results = myList

	}()

	//h5 analysis
	wg.Add(1)
	go func() {
		defer wg.Done()
		myMap := h5Analysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["h5"]
		mu.RUnlock()
		ruleResults.H5Results = myList

	}()

	//h6 analysis
	wg.Add(1)
	go func() {
		defer wg.Done()
		myMap := h6Analysis(l.l, nodeMap)
		mu.RLock()
		myList := myMap["h6"]
		mu.RUnlock()
		ruleResults.H6Results = myList

	}()

	wg.Wait()
	return ruleResults
}

//divAnalysis function initiates all the div rule based analysis
func divAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Divtag {
	l.Println("Initiating div tag Analysis......")
	nodes := nodeMap["divNodes"]

	var issues = make(map[string][]Divtag)
	var list []Divtag
	for _, node := range nodes {
		var tag Divtag

		//build the node
		tag.Div = nodeText(node)

		//implement rules
		tag.divRulesWCAG111(node, l)

		//l.Printf("divTag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("divs: %v\n", list)
	issues["div"] = list

	return issues
}

//buttonAnalysis function initiates all the button rule based analysis
func buttonAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Buttontag {
	l.Println("Initiating button tag Analysis......")
	nodes := nodeMap["buttonNodes"]

	var issues = make(map[string][]Buttontag)
	var list []Buttontag
	for _, node := range nodes {
		var tag Buttontag

		//build the node
		tag.Button = nodeText(node)

		//implement rules
		tag.buttonRulesWCAG111(node, l)
		//l.Printf("buttonTag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("buttons: %v\n", list)
	issues["button"] = list

	return issues
}

//inputAnalysis function initiates all the input rule based analysis
func inputAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Inputtag {
	l.Println("Initiating input tag Analysis......")
	nodes := nodeMap["inputNodes"]

	var issues = make(map[string][]Inputtag)
	var list []Inputtag
	for _, node := range nodes {
		var tag Inputtag

		//build the node
		tag.Input = nodeText(node)

		//implement the rules
		tag.inputRulesWCAG111(node, l)
		//l.Printf("inputTag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("inputs: %v\n", list)
	issues["input"] = list

	return issues
}

//imagesAnalysis function initiates all the images rule based analysis
func imagesAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Imgtag {
	l.Println("Initiating image tag Analysis......")
	nodes := nodeMap["imgNodes"]

	var issues = make(map[string][]Imgtag)
	var list []Imgtag
	for _, node := range nodes {
		var tag Imgtag

		//build the nodes
		tag.Img = nodeText(node)

		//implement the rules
		tag.imagesRulesWCAG111(node, l)
		//l.Printf("imageTag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("images: %v\n", list)
	issues["image"] = list

	return issues
}

//videoAnalysis function initiates all the video rule based analysis
func videoAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Videotag {
	l.Println("Initiating video tag Analysis......")
	nodes := nodeMap["videoNodes"]

	var issues = make(map[string][]Videotag)
	var list []Videotag
	for _, node := range nodes {
		var tag Videotag

		//build the node
		tag.Video = nodeText(node)

		//implement the rules
		tag.videoRulesWCAG111(node, l)
		tag.videoRulesWCAG121(node, l)
		//l.Printf("videoTag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("videos: %v\n", list)
	issues["video"] = list

	return issues
}

//audioAnalysis function initiates all the audio rule based analysis
func audioAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Audiotag {
	l.Println("Initiating audio tag Analysis......")
	nodes := nodeMap["audioNodes"]

	var issues = make(map[string][]Audiotag)
	var list []Audiotag
	for _, node := range nodes {
		var tag Audiotag

		//build the node
		tag.Audio = nodeText(node)

		//implement the rules
		tag.audioRulesWCAG111(node, l)
		tag.audioRulesWCAG121(node, l)
		//l.Printf("audioTag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("audios: %v\n", list)
	issues["audio"] = list

	return issues
}

//textareaAnalysis function initiates all the textarea rule based analysis
func textareaAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Textareatag {
	l.Println("Initiating textarea tag Analysis......")
	nodes := nodeMap["textareaNodes"]

	var issues = make(map[string][]Textareatag)
	var list []Textareatag
	for _, node := range nodes {
		var tag Textareatag

		//build the node
		tag.Textarea = nodeText(node)

		//implement the rule
		tag.textAreaRulesWCAG111(node, l)
		//l.Printf("textareaTag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("textarea: %v\n", list)
	issues["textarea"] = list

	return issues
}

//selectAnalysis function initiates all the select rule based analysis
func selectAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Selecttag {
	l.Println("Initiating select tag Analysis......")
	nodes := nodeMap["selectNodes"]

	var issues = make(map[string][]Selecttag)
	var list []Selecttag
	for _, node := range nodes {
		var tag Selecttag

		//build the node
		tag.Select = nodeText(node)

		//implement the rule
		tag.selectRulesWCAG111(node, l)
		//l.Printf("selectTag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("select: %v\n", list)
	issues["select"] = list

	return issues
}

//iframeAnalysis function initiates all the iframe rule based analysis
func iframeAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Iframetag {
	l.Println("Initiating iFrame tag Analysis......")
	nodes := nodeMap["iframeNodes"]

	var issues = make(map[string][]Iframetag)
	var list []Iframetag
	for _, node := range nodes {
		var tag Iframetag

		//build the node
		tag.Iframe = nodeText(node)

		//implement the rule
		tag.iframeRulesWCAG111(node, l)
		//l.Printf("iframeTag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("iframe: %v\n", list)
	issues["iframe"] = list

	return issues
}

//linkAnalysis function initiates all the link rule based analysis
func linkAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Anchortag {
	l.Println("Initiating Anchor tag Analysis......")
	nodes := nodeMap["anchorNodes"]

	var issues = make(map[string][]Anchortag)
	var list []Anchortag
	for _, node := range nodes {
		var tag Anchortag

		//build the node
		tag.Anchor = nodeText(node)

		//implement the rule
		tag.linkRulesWCAG111(node, l)
		//l.Printf("Anchortag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("anchor: %v\n", list)
	issues["anchor"] = list

	return issues
}

//areaAnalysis function initiates all the link rule based analysis
func areaAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Areatag {
	l.Println("Initiating area tag Analysis......")
	nodes := nodeMap["areaNodes"]

	var issues = make(map[string][]Areatag)
	var list []Areatag
	for _, node := range nodes {
		var tag Areatag

		//build the node
		tag.Area = nodeText(node)

		//implement the rule
		tag.areaRulesWCAG111(node, l)
		//l.Printf("Areatag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("Area: %v\n", list)
	issues["area"] = list

	return issues
}

//objectAnalysis function initiates all the object rule based analysis
func objectAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Objecttag {
	l.Println("Initiating object tag Analysis......")
	nodes := nodeMap["objectNodes"]

	var issues = make(map[string][]Objecttag)
	var list []Objecttag
	for _, node := range nodes {
		var tag Objecttag

		//build the node
		tag.Object = nodeText(node)

		//implement the rules
		tag.objectRulesWCAG121(node, l)
		//l.Printf("Objecttag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("Object: %v\n", list)
	issues["object"] = list

	return issues
}

//embedAnalysis function initiates all the embed rule based analysis
func embedAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Embedtag {
	l.Println("Initiating embed tag Analysis......")
	nodes := nodeMap["embedNodes"]

	var issues = make(map[string][]Embedtag)
	var list []Embedtag
	for _, node := range nodes {
		var tag Embedtag

		//build the node
		tag.Embed = nodeText(node)

		//implement teh rule
		tag.embedRulesWCAG121(node, l)
		//l.Printf("Embedtag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("Embed: %v\n", list)
	issues["embed"] = list

	return issues
}

//trackAnalysis function initiates all the track rule based analysis
func trackAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Tracktag {
	l.Println("Initiating track tag Analysis......")
	nodes := nodeMap["trackNodes"]

	var issues = make(map[string][]Tracktag)
	var list []Tracktag
	for _, node := range nodes {
		var tag Tracktag

		//build the node
		tag.Track = nodeText(node)

		//implement the rule
		tag.trackRulesWCAG121(node, l)
		tag.trackRulesWCAG122(node, l)
		//l.Printf("tracktag : %v", tag)
		list = append(list, tag)
	}
	//l.Printf("track: %v\n", list)
	issues["track"] = list

	return issues
}

//cssAnalysis function initiates all the CSS rule based analysis
func cssAnalysis(l *log.Logger, cssLink []string) {
	l.Println("Initiating track tag Analysis......")
	for _, css := range cssLink {
		l.Printf("CSS : %v ", css)
	}
}

//h1Analysis function initiates all the h1 rule based analysis
func h1Analysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]H1tag {
	l.Println("Initiating h1 tag Analysis......")
	nodes := nodeMap["h1Nodes"]
	var issues = make(map[string][]H1tag)
	var list []H1tag

	for _, node := range nodes {
		var tag H1tag

		//build the node
		tag.H1 = nodeText(node)

		//implement the rule
		tag.h1RulesWCAG241(node, l)
		list = append(list, tag)

	}
	issues["h1"] = list
	return issues
}

//h2Analysis function initiates all the h2 rule based analysis
func h2Analysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]H2tag {
	l.Println("Initiating h2 tag Analysis......")
	nodes := nodeMap["h2Nodes"]
	var issues = make(map[string][]H2tag)
	var list []H2tag

	for _, node := range nodes {
		var tag H2tag

		//build the node
		tag.H2 = nodeText(node)

		//implement the rule
		tag.h2RulesWCAG241(node, l)
		list = append(list, tag)

	}
	issues["h2"] = list
	return issues
}

//h3Analysis function initiates all the h3 rule based analysis
func h3Analysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]H3tag {
	l.Println("Initiating h3 tag Analysis......")
	nodes := nodeMap["h3Nodes"]
	var issues = make(map[string][]H3tag)
	var list []H3tag

	for _, node := range nodes {
		var tag H3tag

		//build the node
		tag.H3 = nodeText(node)

		//implement the rule

		list = append(list, tag)

	}
	issues["h3"] = list
	return issues
}

//h4Analysis function initiates all the h4 rule based analysis
func h4Analysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]H4tag {
	l.Println("Initiating h4 tag Analysis......")
	nodes := nodeMap["h4Nodes"]
	var issues = make(map[string][]H4tag)
	var list []H4tag

	for _, node := range nodes {
		var tag H4tag

		//build the node
		tag.H4 = nodeText(node)

		//implement the rule
		list = append(list, tag)

	}
	issues["h4"] = list
	return issues
}

//h5Analysis function initiates all the h5 rule based analysis
func h5Analysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]H5tag {
	l.Println("Initiating h5 tag Analysis......")
	nodes := nodeMap["h5Nodes"]
	var issues = make(map[string][]H5tag)
	var list []H5tag

	for _, node := range nodes {
		var tag H5tag

		//build the node
		tag.H5 = nodeText(node)

		//implement the rule

		list = append(list, tag)

	}
	issues["h5"] = list
	return issues
}

//h6Analysis function initiates all the h6 rule based analysis
func h6Analysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]H6tag {
	l.Println("Initiating h6 tag Analysis......")
	nodes := nodeMap["h6Nodes"]
	var issues = make(map[string][]H6tag)
	var list []H6tag

	for _, node := range nodes {
		var tag H6tag

		//build the node
		tag.H6 = nodeText(node)

		//implement the rule

		list = append(list, tag)

	}
	issues["h6"] = list
	return issues
}
