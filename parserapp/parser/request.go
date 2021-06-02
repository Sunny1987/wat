package parser

//MyURLReq is the input request from user
type MyURLReq struct {
	URLFromReq string `json:"url"`
	MaxDepth   int    `json:"depth"`
}
