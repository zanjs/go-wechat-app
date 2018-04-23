package auth

import (
	"fmt"
	"github.com/georgehao/wechat/app/http/controllers"
	"github.com/georgehao/wechat/app/middleware"
	"github.com/georgehao/wechat/app/models"
	"github.com/georgehao/wechat/config"
	"github.com/gin-gonic/gin"
	"github.com/nanjishidu/wechat/small"
)

type WxLogin struct {
	Code          string `form:"code" json:"code" binding:"required"`
	EncryptedData string `form:"encrypted_data" json:"encrypted_data" binding:"required"`
	Iv            string `form:"iv" json:"iv" binding:"required"`
}

type LoginController struct {
	controllers.Controller
}

// Login
func (controller *LoginController) Login(c *gin.Context) {
	var json WxLogin
	if err := c.ShouldBindJSON(&json); err != nil {
		controller.Fail(c, config.ErrorWxValidate)
		return
	}

	// 为了测试方便添加跳过微信登录
	if json.Code == "0nQ8QXaK" && json.Iv == "Znp3MyqK" && json.EncryptedData == "Znp94x0O" {
		data := make(map[string]string)
		data["token"] = middleware.CreateToken(fmt.Sprintf("%d", 1))
		controller.Success(c, data)
		return
	}

	wx := small.NewWx(config.WxAppId, config.WxAppSecret)
	// 根据 code 获取用户 session_key 等信息, 返回用户openid 和 session_key
	wxSession, err := wx.GetWxSessionKey(json.Code)
	if err != nil || wxSession.ErrCode != 0 {
		controller.Fail(c, config.ErrorWxLogin)
		return
	}

	// 获取解密后的用户信息
	userInfo, err := small.GetWxUserInfo(wxSession.SessionKey, json.EncryptedData, json.Iv)
	if err != nil {
		controller.Fail(c, config.ErrorWxUserInfo)
		return
	}

	userAttachment := model.UsersAttachments{
		AvatarUrl:  userInfo.AvatarUrl,
		NickName:   userInfo.NickName,
		Gender:     userInfo.Gender,
		SessionKey: wxSession.SessionKey,
	}

	user := model.Users{}
	has, err := config.LukaWechatDB.Where("openid = ? and channel = ?", userInfo.OpenId, model.WX_APP).Get(&user)
	if err != nil {
		controller.Fail(c, config.ErrorGetUser)
		return
	}

	if !has {
		user.OpenId = userInfo.OpenId
		user.Channel = model.WX_APP
		user.UnionId = userInfo.UnionId
		config.LukaWechatDB.Insert(&user)

		userAttachment.UserId = user.Id
		config.LukaWechatDB.Insert(&userAttachment)
	} else {
		config.LukaWechatDB.Where("user_id = ?", user.Id).Update(&userAttachment)
	}

	data := make(map[string]string)
	data["token"] = middleware.CreateToken(fmt.Sprintf("%d", user.Id))
	controller.Success(c, data)
	return
}
