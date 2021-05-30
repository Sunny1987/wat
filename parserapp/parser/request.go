package parser

type MyURLReq struct {
	URLFromReq string `json:"url"`
	MaxDepth   int    `json:"depth"`
}
