package analyzerapp

import (
	"golang.org/x/net/html"
	"strings"
)

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
		if hasChildren(n) || text(n) != "" {
			res = "<div " + str + ">" + text(n) + "</div>"
		} else {
			res = "<div " + str + "/>"
		}

	case "button":
		if hasChildren(n) || text(n) != "" {
			res = "<button " + str + ">" + text(n) + "</button>"
		} else {
			res = "<button " + str + "/>"
		}

	case "input":
		if hasChildren(n) || text(n) != "" {
			res = "<button " + str + ">" + text(n) + "</button>"
		} else {
			res = "<input " + str + "/>"
		}

	case "img":
		res = "<img " + str + "/>"
	case "select":
		res = "<select " + str + "/>"
	case "textarea":
		if hasChildren(n) || text(n) != "" {
			res = "<textarea " + str + ">" + text(n) + "</textarea>"
		} else {
			res = "<textarea " + str + "/>"
		}
	case "a":
		if hasChildren(n) || text(n) != "" {
			res = "<a " + str + ">" + text(n) + "</a>"
		} else {
			res = "<a " + str + "/>"
		}
	case "span":
		if hasChildren(n) || text(n) != "" {
			res = "<span " + str + ">" + text(n) + "</span>"
		} else {
			res = "<span " + str + "/>"
		}
	case "h1":
		res = "<h1 " + str + ">" + text(n) + "</h1>"
	case "h2":
		res = "<h2 " + str + ">" + text(n) + "</h2>"
	case "h3":
		res = "<h3 " + str + ">" + text(n) + "</h3>"
	case "h4":
		res = "<h4 " + str + ">" + text(n) + "</h4>"
	case "h5":
		res = "<h5 " + str + ">" + text(n) + "</h5>"
	case "h6":
		res = "<h6 " + str + ">" + text(n) + "</h6>"
	case "p":
		if hasChildren(n) || text(n) != "" {
			res = "<p " + str + ">" + text(n) + "</p>"
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
	if n.Type == html.TextNode {
		return true
	}
	return false
}

//text will get the text value of the node
func text(n *html.Node) string {
	if isTextNode(n) {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = text(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}
