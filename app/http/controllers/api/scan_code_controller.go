package api

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/georgehao/wechat/app/http/controllers"
	"github.com/georgehao/wechat/app/models"
	"github.com/georgehao/wechat/app/schemas"
	"github.com/georgehao/wechat/config"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type ScanCodeController struct {
	controllers.Controller
}

type PictureBook struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Isbn   string `form:"isbn" json:"isbn" binding:"required"`
	Status int    `form:"status" json:"status"`
}

// Scan 扫码
func (controller *ScanCodeController) Scan(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetString("userID"))
	if err != nil {
		controller.Fail(c, config.ErrorAuth)
		return
	}

	attachment := model.UsersAttachments{}
	has, err := config.LukaWechatDB.Where("user_id = ?", userId).Get(&attachment)
	if err != nil || !has {
		controller.Fail(c, config.ErrorGetUser)
		return
	}

	record := model.ScanRecords{}
	// if user isn't vip, he can only scan total 500; and can scan 200/day
	if attachment.IsVIP == 0 {
		total, err := config.LukaWechatDB.Where("user_id = ?", userId).Count(&record)
		if err != nil {
			controller.Fail(c, config.ErrorUserRecord)
			return
		}

		if total >= config.TotalScanRecords {
			controller.Fail(c, config.ErrorTotalScan)
			return
		}

		//get current YY-MM-DD
		var date string = time.Now().Format("2006-01-02")
		if config.Env == "production" {
			h, _ := time.ParseDuration("1h")
			date = time.Now().Add(15 * h).Format("2006-01-02")
		}

		date = "%" + date + "%"
		dayScans, err := config.LukaWechatDB.Where("user_id = ? and created_at LIKE ?", userId, date).Count(&record)
		if dayScans >= config.DayScanRecords {
			controller.Fail(c, config.ErrorDayScan)
			return
		}
	}

	// 先从picture_books查找数据, 如果查不到再去第三方查找
	books := []PictureBook{}
	isbn := c.Param("isbn")
	key := fmt.Sprintf("wechat:picture_book:%s", isbn)
	results, err := redis.Strings(config.Redis.Do("LRANGE", key, 0, 100))
	if len(results) > 0 {
		for _, result := range results {
			fields := strings.SplitN(result, ":", 3)
			status, _ := strconv.Atoi(fields[2])
			pictureBook := PictureBook{
				Name:   fields[0],
				Isbn:   fields[1],
				Status: status,
			}
			books = append(books, pictureBook)
		}
	} else {
		key := fmt.Sprintf("wechat:third_book:%s", isbn)
		results, _ := redis.Strings(config.Redis.Do("LRANGE", key, 0, 100))
		if len(results) > 0 {
			for _, result := range results {
				fields := strings.SplitN(result, ":", 3)
				status, _ := strconv.Atoi(fields[2])
				pictureBook := PictureBook{
					Name:   fields[0],
					Isbn:   fields[1],
					Status: status,
				}
				books = append(books, pictureBook)
			}
		} else {
			pictureBook := PictureBook{
				Name:   "未知",
				Isbn:   "未知",
				Status: 0,
			}
			books = append(books, pictureBook)
		}

	}

	if len(books) == 1 {
		scanRecord := model.ScanRecords{
			UserId: userId,
			Name:   books[0].Name,
			Isbn:   books[0].Isbn,
			Status: books[0].Status,
		}
		config.LukaWechatDB.Insert(&scanRecord)
	}

	controller.Success(c, books)
}

// Store 入库扫码记录
func (controller *ScanCodeController) Store(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetString("userID"))
	if err != nil {
		controller.Fail(c, config.ErrorAuth)
		return
	}

	var json PictureBook
	if err := c.ShouldBindJSON(&json); err != nil {
		controller.Fail(c, config.ErrorValidate)
		return
	}

	scanRecord := model.ScanRecords{
		UserId: userId,
		Name:   json.Name,
		Isbn:   json.Isbn,
		Status: json.Status,
	}
	config.LukaWechatDB.Insert(&scanRecord)

	controller.Success(c, nil)
}

// Index
func (controller *ScanCodeController) Index(c *gin.Context) {
	userId, err := strconv.Atoi(c.GetString("userID"))
	if err != nil {
		controller.Fail(c, config.ErrorAuth)
		return
	}

	pageString := c.DefaultQuery("page", "0")
	sizeString := c.DefaultQuery("size", "10")

	page, _ := strconv.Atoi(pageString)
	size, _ := strconv.Atoi(sizeString)

	page = page - 1
	if page < 0 {
		page = 0
	}

	scanRecords := make([]schemas.ScanRecords, 0)
	err = config.LukaWechatDB.
		Select("scan_records.id, scan_records.name, scan_records.isbn, scan_records.status, scan_records.updated_at").
		Where("scan_records.user_id = ?", userId).
		Limit(size, page*size).
		Desc("id").
		Find(&scanRecords)

	if err != nil {
		controller.Fail(c, config.ErrorUserRecord)
	}
	controller.Success(c, scanRecords)
}
