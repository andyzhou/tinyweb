package json

/*
 * interface of json
 */
type IJson interface {
	EncodeSelf() ([]byte, error)
	DecodeSelf(data []byte) error
	Encode(i interface{}) ([]byte, error)
	Decode(data []byte, val interface{}) error
	EncodeSimple(data map[string]interface{}) ([]byte, error)
	DecodeSimple(data []byte, kv map[string]interface{}) error
}