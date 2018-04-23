package middleware

import (
	"github.com/appleboy/gin-jwt"
	"github.com/georgehao/wechat/app/http/controllers"
	"github.com/georgehao/wechat/app/models"
	"github.com/georgehao/wechat/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var AuthMiddlerware jwt.GinJWTMiddleware

// CreateToken
func CreateToken(userId string) string {
	AuthMiddlerware.MiddlewareInit()
	return AuthMiddlerware.TokenGenerator(userId)
}

func authorizator(userId string, c *gin.Context) bool {
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic("convert string to int error")
	}

	user := model.Users{}
	has, err := config.LukaWechatDB.Id(id).Get(&user)
	if err != nil {
		return false
	}

	if has {
		return true
	}

	return false
}

func init() {
	AuthMiddlerware = jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte(config.JwtSecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return userId, true
			}
			return userId, false
		},
		Authorizator: authorizator,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.IndentedJSON(http.StatusOK, controllers.Unauthorized(message))
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}
