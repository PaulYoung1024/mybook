package controllers

import "mybook/models"

type ExploreController struct {
	BaseController
}

func (c *ExploreController) Index() {
	var (
		cid  int
		cate models.Category
	)
	if cid, _ = c.GetInt("cid"); cid > 0 {
		c.Data["Cid"] = cid
		cate = new(models.Category).Find(cid)
		c.Data["Cate"] = cate
	}
	c.TplName = "explore/index.html"

	pageIndex, _ := c.GetInt("page", 1)
	pageSize := 24

}
