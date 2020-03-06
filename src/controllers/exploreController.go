package controllers

import (
	"math"
	"mybook/models"
	"mybook/utils"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/logs"
)

type ExploreController struct {
	BaseController
}

func (c *ExploreController) Index() {
	var (
		cid       int
		cate      models.Category
		urlPrefix = beego.URLFor("ExploreController.Index")
	)
	if cid, _ = c.GetInt("cid"); cid > 0 {
		c.Data["Cid"] = cid
		cate = new(models.Category).Find(cid)
		c.Data["Cate"] = cate
	}
	c.TplName = "explore/index.html"

	pageIndex, _ := c.GetInt("page", 1)
	pageSize := 24

	books, totalCount, err := new(models.Book).HomeData(pageIndex, pageSize, cid)
	if err != nil {
		logs.Error(err)
		c.Abort("404")
	}

	if totalCount > 0 {
		urlSuffix := ""
		if cid > 0 {
			urlSuffix = urlSuffix + "&cid=" + strconv.Itoa(cid)
		}
		html := utils.NewPaginations(4, totalCount, pageSize, pageIndex, urlPrefix, urlSuffix)
		c.Data["PageHtml"] = html
	} else {
		c.Data["PageHtml"] = ""
	}
	c.Data["TotalPages"] = int(math.Ceil(float64(totalCount) / float64(pageSize)))
	c.Data["Lists"] = books
}
