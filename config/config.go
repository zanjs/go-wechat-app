package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"strconv"
)

// JwtSecretKey jwt secret key
var JwtSecretKey string

// Env runtime environment
var Env string

// Wechat AppId
var WxAppId string

// Wechat App Secret
var WxAppSecret string

// Mysql db
var LukaWechatDB *xorm.Engine

// redis
var Redis redis.Conn
var RedisPrefix string

// Luka_wechat
var lukaWechat map[string]string
var lukaWechatRedis map[string]string

// debug
var debug bool

// port
var Port int

// Scan
var TotalScanRecords int64
var DayScanRecords int64

// Load env.yaml config and parse config.
func Load() {
	loadLukaWechat()
	LoadRedis()
}

func loadLukaWechat() {
	if _, ok := lukaWechat["connection"]; !ok {
		panic("Luka_chat connection config error")
	}

	if _, ok := lukaWechat["host"]; !ok {
		panic("Luka_chat host config error")
	}

	if _, ok := lukaWechat["port"]; !ok {
		panic("Luka_chat port config error")
	}

	if _, ok := lukaWechat["database"]; !ok {
		panic("Luka_chat database config error")
	}

	if _, ok := lukaWechat["user_name"]; !ok {
		panic("Luka_chat user name config error")
	}

	if _, ok := lukaWechat["password"]; !ok {
		panic("Luka_chat password config error")
	}

	//db:dbadmin@tcp(127.0.0.1:3306)/foo?charset=utf8&parseTime=true&loc=Local
	param := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		lukaWechat["user_name"], lukaWechat["password"], lukaWechat["host"], lukaWechat["port"], lukaWechat["database"])
	//param := "root:123456@/luka_wechat?charset=utf8&parseTime=True&loc=Local"
	var err error
	LukaWechatDB, err = xorm.NewEngine(lukaWechat["connection"], param)
	if err != nil {
		panic("Open xorm failure")
	}

	LukaWechatDB.SetMaxIdleConns(0)
	if debug {
		LukaWechatDB.ShowSQL(true)
		LukaWechatDB.Logger().SetLevel(core.LOG_DEBUG)
	}
}

func LoadRedis() {
	if _, ok := lukaWechatRedis["host"]; !ok {
		panic("Luka_chat redis host config error")
	}

	if _, ok := lukaWechatRedis["password"]; !ok {
		panic("Luka_chat redis password config error")
	}

	if _, ok := lukaWechatRedis["port"]; !ok {
		panic("Luka_chat redis port config error")
	}

	if _, ok := lukaWechatRedis["db"]; !ok {
		panic("Luka_chat redis db config error")
	}

	if _, ok := lukaWechatRedis["prefix"]; !ok {
		panic("Luka_chat redis prefix config error")
	}

	param := fmt.Sprintf("%s:%s", lukaWechatRedis["host"], lukaWechatRedis["port"])
	option := redis.DialPassword(lukaWechatRedis["password"])
	var err error
	Redis, err = redis.Dial("tcp", param, option)
	if err != nil {
		panic("Open redis failure")
	}

	dbInt, _ := strconv.Atoi(lukaWechatRedis["db"])
	Redis.Do("SELECT", dbInt)

	RedisPrefix = lukaWechatRedis["prefix"]
}

func init() {
	viper.SetConfigName("env")  // name of config file (without extension)
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Set model
	model := gin.ReleaseMode
	debug = false
	if viper.GetBool("app.debug") == true {
		model = gin.DebugMode
		debug = true
	}
	gin.SetMode(model)

	Env = viper.GetString("app.env")

	Port = viper.GetInt("app.port")

	JwtSecretKey = viper.GetString("jwt.secret_key")

	WxAppId = viper.GetString("wechat.app_id")

	WxAppSecret = viper.GetString("wechat.app_secret")

	lukaWechat = viper.GetStringMapString("db.luka_wechat")

	TotalScanRecords = viper.GetInt64("scan.total")
	DayScanRecords = viper.GetInt64("scan.day")

	lukaWechatRedis = viper.GetStringMapString("redis")
}
