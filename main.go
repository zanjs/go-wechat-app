package main // simulate some private data

import (
	"fmt"
	"github.com/georgehao/wechat/config"
	"github.com/georgehao/wechat/routes"
	"net/http"
)

func main() {
	config.Load()
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), routes.Engine())
}
