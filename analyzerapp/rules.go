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

	wg.Add(23)
	go ruleResults.divAnalysis(l.l, nodeMap)
	go ruleResults.buttonAnalysis(l.l, nodeMap)
	go ruleResults.inputAnalysis(l.l, nodeMap)
	go ruleResults.imagesAnalysis(l.l, nodeMap)
	go ruleResults.videoAnalysis(l.l, nodeMap)
	go ruleResults.audioAnalysis(l.l, nodeMap)
	go ruleResults.textareaAnalysis(l.l, nodeMap)
	go ruleResults.selectAnalysis(l.l, nodeMap)
	go ruleResults.iframeAnalysis(l.l, nodeMap)
	go ruleResults.linkAnalysis(l.l, nodeMap)
	go ruleResults.areaAnalysis(l.l, nodeMap)
	go ruleResults.objectAnalysis(l.l, nodeMap)
	go ruleResults.embedAnalysis(l.l, nodeMap)
	go ruleResults.trackAnalysis(l.l, nodeMap)
	go ruleResults.h1Analysis(l.l, nodeMap)
	go ruleResults.h2Analysis(l.l, nodeMap)
	go ruleResults.h3Analysis(l.l, nodeMap)
	go ruleResults.h4Analysis(l.l, nodeMap)
	go ruleResults.h5Analysis(l.l, nodeMap)
	go ruleResults.h6Analysis(l.l, nodeMap)
	go ruleResults.paraAnalysis(l.l, nodeMap)
	go ruleResults.preAnalysis(l.l, nodeMap)
	go ruleResults.cssAnalysis(l.l, cssList, nodeMap)

	wg.Wait()
	//cssList = nil
	//nodeMap = nil
	return ruleResults
}

//divAnalysis function initiates all the div rule based analysis
func (resp *Response) divAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating div tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["divNodes"]

	var list []Divtag
	for _, node := range nodes {
		var tag Divtag

		//build the node
		tag.Div = nodeText(node)

		//implement rules
		tag.divRulesWCAG111(node, l)

		list = append(list, tag)
	}
	resp.DivResults = list

}

//buttonAnalysis function initiates all the button rule based analysis
func (resp *Response) buttonAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating button tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["buttonNodes"]

	var list []Buttontag
	for _, node := range nodes {
		var tag Buttontag

		//build the node
		tag.Button = nodeText(node)

		//implement rules
		tag.buttonRulesWCAG111(node, l)
		list = append(list, tag)
	}
	resp.ButtonResults = list

}

//inputAnalysis function initiates all the input rule based analysis
func (resp *Response) inputAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating input tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["inputNodes"]

	var list []Inputtag
	for _, node := range nodes {
		var tag Inputtag

		//build the node
		tag.Input = nodeText(node)

		//implement the rules
		tag.inputRulesWCAG111(node, l)

		list = append(list, tag)
	}
	resp.InputResults = list
}

//imagesAnalysis function initiates all the images rule based analysis
func (resp *Response) imagesAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating image tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["imgNodes"]

	var list []Imgtag
	for _, node := range nodes {
		var tag Imgtag

		//build the nodes
		tag.Img = nodeText(node)

		//implement the rules
		tag.imagesRulesWCAG111(node, l)

		list = append(list, tag)
	}

	resp.ImageResults = list
}

//videoAnalysis function initiates all the video rule based analysis
func (resp *Response) videoAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating video tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["videoNodes"]

	var list []Videotag
	for _, node := range nodes {
		var tag Videotag

		//build the node
		tag.Video = nodeText(node)

		//implement the rules
		tag.videoRulesWCAG111(node, l)
		tag.videoRulesWCAG121(node, l)

		list = append(list, tag)
	}
	resp.VideoResults = list
}

//audioAnalysis function initiates all the audio rule based analysis
func (resp *Response) audioAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating audio tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["audioNodes"]

	var list []Audiotag
	for _, node := range nodes {
		var tag Audiotag

		//build the node
		tag.Audio = nodeText(node)

		//implement the rules
		tag.audioRulesWCAG111(node, l)
		tag.audioRulesWCAG121(node, l)

		list = append(list, tag)
	}
	resp.AudioResults = list
}

//textareaAnalysis function initiates all the textarea rule based analysis
func (resp *Response) textareaAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating textarea tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["textareaNodes"]

	var list []Textareatag
	for _, node := range nodes {
		var tag Textareatag

		//build the node
		tag.Textarea = nodeText(node)

		//implement the rule
		tag.textAreaRulesWCAG111(node, l)

		list = append(list, tag)
	}
	resp.TextareaResults = list
}

//selectAnalysis function initiates all the select rule based analysis
func (resp *Response) selectAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating select tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["selectNodes"]

	var list []Selecttag
	for _, node := range nodes {
		var tag Selecttag

		//build the node
		tag.Select = nodeText(node)

		//implement the rule
		tag.selectRulesWCAG111(node, l)

		list = append(list, tag)
	}
	resp.SelectResults = list
}

//iframeAnalysis function initiates all the iframe rule based analysis
func (resp *Response) iframeAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating iFrame tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["iframeNodes"]

	var list []Iframetag
	for _, node := range nodes {
		var tag Iframetag

		//build the node
		tag.Iframe = nodeText(node)

		//implement the rule
		tag.iframeRulesWCAG111(node, l)

		list = append(list, tag)
	}
	resp.IframeResults = list
}

//linkAnalysis function initiates all the link rule based analysis
func (resp *Response) linkAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating Anchor tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["anchorNodes"]

	var list []Anchortag
	for _, node := range nodes {
		var tag Anchortag

		//build the node
		tag.Anchor = nodeText(node)

		//implement the rule
		tag.linkRulesWCAG111(node, l)

		list = append(list, tag)
	}
	resp.LinkResults = list
}

//areaAnalysis function initiates all the link rule based analysis
func (resp *Response) areaAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating area tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["areaNodes"]

	var list []Areatag
	for _, node := range nodes {
		var tag Areatag

		//build the node
		tag.Area = nodeText(node)

		//implement the rule
		tag.areaRulesWCAG111(node, l)

		list = append(list, tag)
	}
	resp.AreaResults = list
}

//objectAnalysis function initiates all the object rule based analysis
func (resp *Response) objectAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating object tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["objectNodes"]

	var list []Objecttag
	for _, node := range nodes {
		var tag Objecttag

		//build the node
		tag.Object = nodeText(node)

		//implement the rules
		tag.objectRulesWCAG121(node, l)

		list = append(list, tag)
	}
	resp.ObjectResults = list
}

//embedAnalysis function initiates all the embed rule based analysis
func (resp *Response) embedAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating embed tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["embedNodes"]

	var list []Embedtag
	for _, node := range nodes {
		var tag Embedtag

		//build the node
		tag.Embed = nodeText(node)

		//implement teh rule
		tag.embedRulesWCAG121(node, l)

		list = append(list, tag)
	}
	resp.EmbedResults = list
}

//trackAnalysis function initiates all the track rule based analysis
func (resp *Response) trackAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating track tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["trackNodes"]

	var list []Tracktag
	for _, node := range nodes {
		var tag Tracktag

		//build the node
		tag.Track = nodeText(node)

		//implement the rule
		tag.trackRulesWCAG121(node, l)
		tag.trackRulesWCAG122(node, l)

		list = append(list, tag)
	}
	resp.TrackResults = list
}

//cssAnalysis function initiates all the CSS rule based analysis
func (resp *Response) cssAnalysis(l *log.Logger, cssLink []string, nodeMap map[string][]*html.Node) {
	l.Println("Initiating track tag Analysis......")
	defer wg.Done()
	var list []CSS
	nodes := nodeMap["divNodes"]
	for _, css := range cssLink {
		l.Printf("CSS : %v ", css)
		var tag CSS

		//add css data
		tag.CSS = css

		//implement rule
		tag.cssAnalysisWCAG111(css, l, nodes)

	}

	resp.CSSResults = list
}

//h1Analysis function initiates all the h1 rule based analysis
func (resp *Response) h1Analysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating h1 tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["h1Nodes"]
	var list []H1tag

	for _, node := range nodes {
		var tag H1tag

		//build the node
		tag.H1 = nodeText(node)

		//implement the rule
		tag.h1RulesWCAG241(node, l)
		list = append(list, tag)

	}
	resp.H1Results = list
}

//h2Analysis function initiates all the h2 rule based analysis
func (resp *Response) h2Analysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating h2 tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["h2Nodes"]

	var list []H2tag

	for _, node := range nodes {
		var tag H2tag

		//build the node
		tag.H2 = nodeText(node)

		//implement the rule
		tag.h2RulesWCAG241(node, l)
		list = append(list, tag)

	}
	resp.H2Results = list
}

//h3Analysis function initiates all the h3 rule based analysis
func (resp *Response) h3Analysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating h3 tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["h3Nodes"]

	var list []H3tag

	for _, node := range nodes {
		var tag H3tag

		//build the node
		tag.H3 = nodeText(node)

		//implement the rule

		list = append(list, tag)

	}
	resp.H3Results = list
}

//h4Analysis function initiates all the h4 rule based analysis
func (resp *Response) h4Analysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating h4 tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["h4Nodes"]

	var list []H4tag

	for _, node := range nodes {
		var tag H4tag

		//build the node
		tag.H4 = nodeText(node)

		//implement the rule
		list = append(list, tag)

	}
	resp.H4Results = list
}

//h5Analysis function initiates all the h5 rule based analysis
func (resp *Response) h5Analysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating h5 tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["h5Nodes"]
	var list []H5tag

	for _, node := range nodes {
		var tag H5tag

		//build the node
		tag.H5 = nodeText(node)

		//implement the rule

		list = append(list, tag)

	}
	resp.H5Results = list
}

//h6Analysis function initiates all the h6 rule based analysis
func (resp *Response) h6Analysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating h6 tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["h6Nodes"]
	var list []H6tag

	for _, node := range nodes {
		var tag H6tag

		//build the node
		tag.H6 = nodeText(node)

		//implement the rule

		list = append(list, tag)

	}
	resp.H6Results = list
}

//paraAnalysis function initiates all the paragraphs rule based analysis
func (resp *Response) paraAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating para tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["paraNodes"]
	var list []Paratag

	for _, node := range nodes {
		var tag Paratag

		//build the node
		tag.Para = nodeText(node)

		//implement the rule

		list = append(list, tag)

	}
	resp.ParaResults = list
}

//preAnalysis function initiates all the pre rule based analysis
func (resp *Response) preAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) {
	l.Println("Initiating pre tag Analysis......")
	defer wg.Done()
	nodes := nodeMap["preNodes"]
	var list []Pretag

	for _, node := range nodes {
		var tag Pretag

		//build the node
		tag.Pre = nodeText(node)

		//implement the rule
		tag.preRulesWCAG111(node, l)

		list = append(list, tag)

	}
	resp.PreResults = list
}
