package face

import "reflect"

type Util struct {

}

func (f *Util) ResetJsonObject(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}