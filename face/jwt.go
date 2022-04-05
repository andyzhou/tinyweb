package face

import (
	"errors"
	"github.com/andyzhou/tinyweb/define"
	"github.com/dgrijalva/jwt-go"
	"sync"
)

/*
 * face for jwt encoder/decoder
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//global variable for single instance
var (
	_jwt *Jwt
	_jwtOnce sync.Once
)

//face info
type Jwt struct {
	secret string `secret key string`
	token *jwt.Token `jwt token instance`
	claims jwt.MapClaims `jwt claims object`
}

//get single instance
func GetJwt() *Jwt {
	_jwtOnce.Do(func() {
		_jwt = NewJwt()
	})
	return _jwt
}

//construct
func NewJwt() *Jwt {
	this := &Jwt{
		secret: define.SecretKeyOfJwt,
		token:jwt.New(jwt.SigningMethodHS256),
		claims:make(jwt.MapClaims),
	}
	return this
}

//encode
func (j *Jwt) Encode(input map[string]interface{}) (string, error) {
	j.claims = input
	j.token.Claims = j.claims
	result, err := j.token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return result, nil
}

//decode
func (j *Jwt) Decode(input string) (map[string]interface{}, error) {
	//parse input string
	token, err := jwt.Parse(input, j.getValidationKey)
	if err != nil {
		return nil, err
	}
	//check header
	if jwt.SigningMethodHS256.Alg() != token.Header["alg"] {
		return nil, errors.New("header error")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//set security key
func (j *Jwt) SetSecurityKey(secret string) {
	if secret == "" {
		return
	}
	j.secret = secret
	return
}

//get validate key
func (j *Jwt) getValidationKey(*jwt.Token) (interface{}, error) {
	return []byte(j.secret), nil
}
