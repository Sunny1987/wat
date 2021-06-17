package parserhandler

//MyURLReq is the input request from user
type MyURLReq struct {
	URLFromReq string `json:"url" validate:"required,url"`
	MaxDepth   int    `json:"depth" validate:"required,depth"`
}
