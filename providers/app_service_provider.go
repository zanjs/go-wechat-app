package providers

import (
	"github.com/georgehao/wechat/app/http/controllers/api"
	"github.com/georgehao/wechat/app/http/controllers/auth"
)

var LoginController *auth.LoginController

var ScanCodeController *api.ScanCodeController

func init() {
	LoginController = new(auth.LoginController)
	ScanCodeController = new(api.ScanCodeController)
}
