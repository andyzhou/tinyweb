package iface

import "github.com/kataras/iris/v12/view"

type ITpl interface {
	RegisterTplFunc(tpl *view.HTMLEngine)
}
