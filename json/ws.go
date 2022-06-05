package json
/*
 * json for web socket
 */

//header
type WSHeader struct {
	Module string `json:"module"`
	Sub string `json:"sub"`
}

//common request
type WSCommonReq struct {
	Header *WSHeader `json:"header"`
}

//net base
type NetBaseInfo struct {
	Session string `json:"session"`
	UserId string `json:"userId"`
	ClientIPv4 string `json:"clientIPv4"`
	ContentType string `json:"contentType"`
	Platform string `json:"platform"`
	Version string `json:"version"`
	ReqTime int64 `json:"reqTime"`
}