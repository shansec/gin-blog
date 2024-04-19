package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"myginblog/utils"
	"myginblog/utils/errMsg"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type LoginClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// SetToken 生成Token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(5 * time.Hour)
	SetClaims := LoginClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "myginblog",
		},
	}
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", errMsg.ERROR
	}
	return token, errMsg.SUCCESS
}

// CheckToken 验证Token
func CheckToken(token string) (*LoginClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*LoginClaims); setToken.Valid {
		return key, errMsg.SUCCESS
	} else {
		return nil, errMsg.ERROR
	}
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errMsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errMsg.GetCodeMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errMsg.GetCodeMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errMsg.GetCodeMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errMsg.ERROR {
			code = errMsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errMsg.GetCodeMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
