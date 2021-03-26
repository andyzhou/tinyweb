package iface

//sign conf
type SignConf struct {
	Switcher bool
	SignKey string
	SkipReqPara []string
}

type ISignature interface {
	GenSign(fields map[string]string) (bool, string)
	AddSkipFields(fields []string) bool
	SetConf(conf *SignConf) bool
	GetSwitcher() bool
}
