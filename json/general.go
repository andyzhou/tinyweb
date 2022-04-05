package json

/*
 * json for general data
 */

//general slice json
type GenSliceJson struct {
	Data []interface{} `json:"data"`
	BaseJson
}

//construct
func NewGenSliceJson() *GenSliceJson {
	this := &GenSliceJson{
		Data: make([]interface{}, 0),
	}
	return this
}

func (j *GenSliceJson) AddData(data ...interface{}) bool {
	if data == nil {
		return false
	}
	j.Data = append(j.Data, data...)
	return false
}