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
