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
		res = "<div " + str + ">" + n.FirstChild.Data + "</div>"
	case "button":
		res = "<button " + str + ">" + n.FirstChild.Data + "</button>"
	case "input":
		res = "<input " + str + ">" + n.FirstChild.Data + "</input>"
	case "img":
		res = "<img " + str + ">" + "/>"
	case "select":
		res = "<select " + str + "/>"
	case "textarea":
		res = "<textarea " + str + ">" + n.FirstChild.Data + "</textarea>"
	case "a":
		res = "<a " + str + ">" + n.FirstChild.Data + "</a>"
	case "span":
		res = "<span " + str + ">" + n.FirstChild.Data + "</span>"
	case "h1":
		res = "<h1 " + str + ">" + n.FirstChild.Data + "</h1>"
	case "h2":
		res = "<h2 " + str + ">" + n.FirstChild.Data + "</h2>"
	case "h3":
		res = "<h3 " + str + ">" + n.FirstChild.Data + "</h3>"
	case "h4":
		res = "<h4 " + str + ">" + n.FirstChild.Data + "</h4>"
	case "h5":
		res = "<h5 " + str + ">" + n.FirstChild.Data + "</h5>"
	case "h6":
		res = "<h6 " + str + ">" + n.FirstChild.Data + "</h6>"
	case "p":
		res = "<p " + str + ">" + n.FirstChild.Data + "</p>"
	case "video":
		res = "<video " + str + "/>"
	case "audio":
		res = "<audio " + str + "/>"

	}

	return res
}
