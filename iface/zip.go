package iface

type IZip interface {
	UnZipData(resp []byte) (bool, map[string]interface{})
	ZipData(req map[string]interface{}) (bool, []byte)
}
