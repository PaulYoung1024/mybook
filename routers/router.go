package routers

import (
	"mybook/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	//首页&分类&详情
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/explore/:cid", &controllers.ExploreController{}, "*:Index")
}
