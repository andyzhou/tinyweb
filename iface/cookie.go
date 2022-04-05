package iface

import (
	"github.com/gin-gonic/gin"
)

type ICookie interface {
	DelCookie(key, domain string, c *gin.Context) error
	GetCookie(key string, c *gin.Context) (map[string]interface{}, error)
	SetCookie(
			key string,
			val map[string]interface{},
			expireSeconds int,
			domain string,
			c *gin.Context,
		) error
	SetJwt(jwt IJwt) bool
}
