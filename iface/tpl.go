package iface

import "html/template"

type ITpl interface {
	GetFuncMap() template.FuncMap
	RegisterTplFunc(name string, cb interface{}) bool
	RegisterTplBaseFunc() bool
}
