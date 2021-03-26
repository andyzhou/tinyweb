package face

import (
	"errors"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/schema"
	"log"
	"net/http"
)

/*
 * face for rest service
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//face info
type Rest struct {
	httpPort int
	ws *restful.WebService
	decoder *schema.Decoder
}

//construct
func NewRest(httpPort int) *Rest {
	//self init
	this := &Rest{
		httpPort:httpPort,
		ws:new(restful.WebService),
		decoder:schema.NewDecoder(),
	}
	//inter init
	this.interInit()
	return this
}

//stop service
func (s *Rest) Stop() {
	if s.ws == nil {
		return
	}
}

//set service
func (s *Rest) Start() {
	//add default container
	restful.DefaultContainer.Add(s.ws)

	log.Println("start web service on port", s.httpPort)
	//start http service
	portStr := fmt.Sprintf(":%d", s.httpPort)
	err := http.ListenAndServe(portStr, nil)
	if err != nil {
		log.Println("Rest start failed, err:", err.Error())
		panic(err)
	}
}

//generate response json data
func (s *Rest) GenJsonResp(jsonObj interface{}, resp *restful.Response) {
	if jsonObj == nil || resp == nil {
		return
	}
	//send to client side
	err := resp.WriteAsJson(jsonObj)
	if err != nil {
		log.Println("Rest::GenJsonResp failed, err:", err.Error())
	}
}


//parse request form
func (s *Rest) ParseReqForm(
					formFace interface{},
					req *restful.Request,
				) error {
	//basic check
	if formFace == nil || req == nil {
		return errors.New("invalid parameters")
	}

	//parse post form
	err := req.Request.ParseForm()
	if err != nil {
		return err
	}

	//decode form data
	err = s.decoder.Decode(formFace, req.Request.PostForm)
	if err != nil {
		return err
	}
	return nil
}


//register dynamic sub route
//dynamicRootUrl like /test/{para1}/{para2}
//should use request.PathParameter("para1") to get real path parameter value
func (s *Rest) RegisterDynamicSubRoute(
					method string,
					consumes string,
					dynamicRootUrl string,
					dynamicPathSlice []string,
					routeFunc restful.RouteFunction,
				) bool {

	//basic check
	if dynamicRootUrl == "" || routeFunc == nil || s.ws == nil {
		return false
	}

	//init new route builder
	rb := new(restful.RouteBuilder)
	rb.Method(method).Path(dynamicRootUrl).To(routeFunc)

	if consumes != "" {
		rb.Consumes(consumes)
	}

	//init path parameter
	if dynamicPathSlice != nil && len(dynamicPathSlice) > 0 {
		for _, key := range dynamicPathSlice {
			pp := s.CreatePathParameter(key, "string")
			rb.Param(pp)
		}
	}

	//add sub route
	s.ws.Route(rb)

	return true
}

//register static sub route
func (s *Rest) RegisterSubRoute(
					method, routeUrl, consumes string,
					parameters [] *restful.Parameter,
					routeFunc restful.RouteFunction,
				) bool {

	//basic check
	if method == "" || routeUrl == "" {
		return false
	}
	if routeFunc == nil || s.ws == nil {
		return false
	}

	//init new route builder
	rb := new(restful.RouteBuilder)

	//set method, request url and route func
	rb.Method(method).Path(routeUrl).To(routeFunc)

	if consumes != "" {
		rb.Consumes(consumes)
	}

	//set parameter
	if parameters != nil && len(parameters) > 0 {
		for _, parameter := range parameters {
			//set sub parameter
			rb.Param(parameter)
		}
	}

	//add sub route
	s.ws.Route(rb)

	//ws.PathParameter("key", "parameter for key").DataType("string")

	return true
}

//create ws form parameter
func (s *Rest) CreateParameter(
					name, kind, defaultVal string,
				) *restful.Parameter {
	//basic check
	if name == "" || kind == "" {
		return nil
	}
	//init new
	param := s.ws.FormParameter(name, "").DataType(kind).DefaultValue(defaultVal)
	return param
}

//create ws path parameter
func (s *Rest) CreatePathParameter(
					name, kind string,
				) *restful.Parameter {
	//basic check
	if name == "" || kind == "" {
		return nil
	}
	//init new
	param := s.ws.PathParameter(name, "").DataType(kind)
	return param
}

////////////////
//private func
////////////////

//inter init
func (s *Rest) interInit() {
	if s.httpPort <= 0 {
		return
	}
	//set mime kind, use json format
	s.ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
}