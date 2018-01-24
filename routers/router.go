package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/blogs", &controllers.BlogCollectionController{})
	beego.Router("/blogs/:id", &controllers.BlogController{})
	beego.Router("/blogs/:id/edit", &controllers.BlogEditController{})
	beego.Router("/blogs/:id/delete", &controllers.BlogDeleteController{})
	beego.Router("/blogs/create", &controllers.BlogCreateController{})
}
