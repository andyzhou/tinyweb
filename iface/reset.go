package iface

import (
	"github.com/emicklei/go-restful"
)

type IReset interface {
	Start()
	GenJsonResp(jsonObj interface{}, resp *restful.Response)
	ParseReqForm(formFace interface{}, req *restful.Request) error

	RegisterDynamicSubRoute(
		method string,
		consumes string,
		dynamicRootUrl string,
		dynamicPathSlice []string,
		routeFunc restful.RouteFunction,
	) bool

	RegisterSubRoute(
		method, routeUrl, consumes string,
		parameters [] *restful.Parameter,
		routeFunc restful.RouteFunction,
	) bool

	CreateParameter(
		name, kind, defaultVal string,
	) *restful.Parameter

	CreatePathParameter(
		name, kind string,
	) *restful.Parameter
}
