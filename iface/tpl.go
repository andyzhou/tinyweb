package iface

type ITpl interface {
	RegisterTplFunc(name string, cb interface{}) bool
	RegisterTplBaseFunc() bool
}
