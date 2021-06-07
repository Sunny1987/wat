package analyzerapp

import "golang.org/x/net/html"

//attributeSearch will respond if the attribute is present
func attributeSearch(attrList []html.Attribute, key string) bool {
	for _, att := range attrList {
		if att.Key == key {
			return true
		}
	}
	return false
}

//attributeCheckVal will respond if the attribute is present and check its value
func attributeCheckVal(attrList []html.Attribute, key string, val string) bool {
	for _, att := range attrList {
		if att.Key == key && att.Val == val {
			return true
		}
	}
	return false
}

//attributeCheckValEmpty will respond if the attribute is present and value is empty
func attributeCheckValEmpty(attrList []html.Attribute, key string) bool {
	for _, att := range attrList {
		if att.Key == key && att.Val != "" {
			return true
		}
	}
	return false
}

//hasNoChild will check if node has child
func hasNoChild(n *html.Node) bool {
	return n.FirstChild == nil
}

//hasChildren will check if node has children
func hasChildren(n *html.Node) bool {
	return n.FirstChild != nil && n.FirstChild != n.LastChild
}

//hasOneChild will check if node has only one child
func hasOneChild(n *html.Node) bool {
	return n.FirstChild != nil && n.FirstChild == n.LastChild
}

//nodeText will retrieve node text
func nodeText(n *html.Node) string {
	var str string
	for _, a := range n.Attr {
		s := a.Key + "=\"" + a.Val + "\" "
		str += s
	}

	var res string

	switch n.Data {
	case "div":
		if hasChildren(n) {
			res = "<div " + str + ">" + n.FirstChild.Data + "</div>"
		} else {
			res = "<div " + str + "/>"
		}

	case "button":
		if hasChildren(n) {
			res = "<button " + str + ">" + n.FirstChild.Data + "</button>"
		} else {
			res = "<button " + str + "/>"
		}

	case "input":
		if hasChildren(n) {
			res = "<button " + str + ">" + n.FirstChild.Data + "</button>"
		} else {
			res = "<input " + str + "/>"
		}

	case "img":
		res = "<img " + str + "/>"
	case "select":
		res = "<select " + str + "/>"
	case "textarea":
		if hasChildren(n) {
			res = "<textarea " + str + ">" + n.FirstChild.Data + "</textarea>"
		} else {
			res = "<textarea " + str + "/>"
		}
	case "a":
		if hasChildren(n) {
			res = "<a " + str + ">" + n.FirstChild.Data + "</a>"
		} else {
			res = "<a " + str + "/>"
		}
	case "span":
		if hasChildren(n) {
			res = "<span " + str + ">" + n.FirstChild.Data + "</span>"
		} else {
			res = "<span " + str + "/>"
		}
	case "h1":
		if hasChildren(n) {
			res = "<h1 " + str + ">" + n.FirstChild.Data + "</h1>"
		} else {
			res = "<h1 " + str + "/>"
		}
	case "h2":
		if hasChildren(n) {
			res = "<h2 " + str + ">" + n.FirstChild.Data + "</h2>"
		} else {
			res = "<h2 " + str + "/>"
		}
	case "h3":
		if hasChildren(n) {
			res = "<h3 " + str + ">" + n.FirstChild.Data + "</h3>"
		} else {
			res = "<h3 " + str + "/>"
		}
	case "h4":
		if hasChildren(n) {
			res = "<h4 " + str + ">" + n.FirstChild.Data + "</h4>"
		} else {
			res = "<h4 " + str + "/>"
		}

	case "h5":
		if hasChildren(n) {
			res = "<h5 " + str + ">" + n.FirstChild.Data + "</h5>"
		} else {
			res = "<h5 " + str + "/>"
		}

	case "h6":
		if hasChildren(n) {
			res = "<h6 " + str + ">" + n.FirstChild.Data + "</h6>"
		} else {
			res = "<h6 " + str + "/>"
		}

	case "p":
		if hasChildren(n) {
			res = "<p " + str + ">" + n.FirstChild.Data + "</p>"
		} else {
			res = "<p " + str + "/>"
		}

	case "video":
		res = "<video " + str + "/>"
	case "audio":
		res = "<audio " + str + "/>"

	}

	return res
}

//isTextNode will validate if node is TextNode
func isTextNode(n *html.Node) bool {
	if n.Type == html.TextNode && n.DataAtom == 0 {
		return true
	}
	return false
}
