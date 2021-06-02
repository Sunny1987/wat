package analyzerapp

import "golang.org/x/net/html"

func attributeSearch(attrList []html.Attribute, key string) bool {
	for _, att := range attrList {
		if att.Key == key {
			return true
		}
	}
	return false
}

func attributeCheckVal(attrList []html.Attribute, key string,val string) bool {
	for _, att := range attrList {
		if att.Key == key && att.Val == val {
			return true
		}
	}
	return false
}

func attributeCheckValEmpty(attrList []html.Attribute, key string) bool {
	for _, att := range attrList {
		if att.Key == key && att.Val != "" {
			return true
		}
	}
	return false
}

func hasNoChild(n *html.Node) bool {
	return n.FirstChild == nil
}

func hasChildren(n *html.Node) bool {
	return n.FirstChild != nil && n.FirstChild != n.LastChild
}

func hasOneChild(n *html.Node) bool {
	return n.FirstChild != nil && n.FirstChild == n.LastChild
}

func nodeText(n *html.Node) string {
	var str string
	for _, a := range n.Attr {
		s := a.Key + "=" + a.Val + " "
		str += s
	}

	var res string

	switch n.Data {
	case "div":
		res = "<div " +str + ">" + n.FirstChild.Data +"</div>"


	}

	return res
}