package iface

type IJwt interface {
	Decode(input string) (map[string]interface{}, error)
	Encode(input map[string]interface{}) (string, error)
	SetSecurityKey(secret string)
}
