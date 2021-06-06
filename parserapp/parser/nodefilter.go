package parser

import "golang.org/x/net/html"

//FilterLinkNodes will give the list of links
func filterLinkNodes(n *html.Node) []*html.Node {

	if isAnchor(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterLinkNodes(c)...)
	}
	return retLnk
}

//filterDivNodes will give the list of divs
func filterDivNodes(n *html.Node) []*html.Node {
	if isDiv(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterDivNodes(c)...)
	}
	return retLnk
}

//filterDivNodes will give the list of divs
func filterParaNodes(n *html.Node) []*html.Node {
	if isParagraph(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterParaNodes(c)...)
	}
	return retLnk
}

//filterSpanNodes will give the list of spans
func filterSpanNodes(n *html.Node) []*html.Node {
	if isSpan(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterSpanNodes(c)...)
	}
	return retLnk
}

//filterH1Nodes will give the list of H1
func filterH1Nodes(n *html.Node) []*html.Node {
	if isH1(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterH1Nodes(c)...)
	}
	return retLnk
}

//filterH2Nodes will give the list of H2
func filterH2Nodes(n *html.Node) []*html.Node {
	if isH2(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterH2Nodes(c)...)
	}
	return retLnk
}

//filterH3Nodes will give the list of H3
func filterH3Nodes(n *html.Node) []*html.Node {
	if isH3(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterH3Nodes(c)...)
	}
	return retLnk
}

//filterH4Nodes will give the list of H4
func filterH4Nodes(n *html.Node) []*html.Node {
	if isH4(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterH4Nodes(c)...)
	}
	return retLnk
}

//filterH5Nodes will give the list of H5
func filterH5Nodes(n *html.Node) []*html.Node {
	if isH5(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterH5Nodes(c)...)
	}
	return retLnk
}

//filterH6Nodes will give the list of divs
func filterH6Nodes(n *html.Node) []*html.Node {
	if isH6(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterH6Nodes(c)...)
	}
	return retLnk
}

//filterImageNodes will give the list of divs
func filterImageNodes(n *html.Node) []*html.Node {
	if isImage(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterImageNodes(c)...)
	}
	return retLnk
}

//filterInputNodes will give the list of divs
func filterInputNodes(n *html.Node) []*html.Node {
	if isInput(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterInputNodes(c)...)
	}
	return retLnk
}

//filterButtonNodes will give the list of divs
func filterButtonNodes(n *html.Node) []*html.Node {
	if isButton(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterButtonNodes(c)...)
	}
	return retLnk
}

//filterVideoNodes will give the list of divs
func filterVideoNodes(n *html.Node) []*html.Node {
	if isVideo(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterVideoNodes(c)...)
	}
	return retLnk
}

//filterAudioNodes will give the list of divs
func filterAudioNodes(n *html.Node) []*html.Node {
	if isAudio(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterAudioNodes(c)...)
	}
	return retLnk
}

//filterSelectNodes will give the list of divs
func filterSelectNodes(n *html.Node) []*html.Node {
	if isSelect(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterSelectNodes(c)...)
	}
	return retLnk
}

//filterTextAreaNodes will give the list of divs
func filterTextAreaNodes(n *html.Node) []*html.Node {
	if isTextArea(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterTextAreaNodes(c)...)
	}
	return retLnk
}

//filterIframeNodes will give the list of divs
func filterIframeNodes(n *html.Node) []*html.Node {
	if isIFrame(n) {
		return []*html.Node{n}
	}
	var retLnk []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retLnk = append(retLnk, filterTextAreaNodes(c)...)
	}
	return retLnk
}
