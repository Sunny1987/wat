package parser

import "golang.org/x/net/html"

//isAnchor will validate an Anchor tag
func isAnchor(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

//isDiv will validate an Div tag
func isDiv(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "div"
}

//isSpan will validate an span tag
func isSpan(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "span"
}

func isH1(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h1"
}

func isH2(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h2"
}

func isH3(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h3"
}

func isH4(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h4"
}

func isH5(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h5"
}

func isH6(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h6"
}

func isParagraph(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "p"
}

func isImage(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "img"
}

func isUL(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "ul"
}

func isOL(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "ol"
}

func isLI(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "li"
}

func isScript(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "script"
}

func isDocTypeOrDocumentType(n *html.Node) bool {
	return n.Type == html.DocumentNode || n.Type == html.DoctypeNode
}

func isTextNode(n *html.Node) bool {
	if n.Type == html.TextNode && n.DataAtom == 0 {
		return true
	}
	return false
}
