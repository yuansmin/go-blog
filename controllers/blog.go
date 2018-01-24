package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/models"
	"strconv"
)

type BlogCollectionController struct {
	beego.Controller
}

func (bc *BlogCollectionController) Get() {
	var blogs []*models.Blog
	o := orm.NewOrm()
	num, err := o.QueryTable("blog").All(&blogs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("query %d blogs", num)
	bc.Data["blogs"] = blogs
	bc.TplName = "blog-list.tpl"
}

type BlogController struct {
	beego.Controller
}

func (bc *BlogController) Get() {
	id, _ := strconv.Atoi(bc.Ctx.Input.Param(":id"))
	blog := &models.Blog{Id: id}
	o := orm.NewOrm()
	err := o.Read(blog)
	bc.Data["blog"] = blog
	bc.TplName = "blog-detail.tpl"
	if err != nil {
		fmt.Println("no such blog~, id=%d", id)
		fmt.Println(err)
		bc.Data["message"] = "not found"
	}
}

type BlogDeleteController struct {
	beego.Controller
}

func (bc *BlogDeleteController) Get() {
	id, _ := strconv.Atoi(bc.Ctx.Input.Param(":id"))
	blog := &models.Blog{Id: id}
	o := orm.NewOrm()
	o.Delete(blog)
	bc.Redirect("/blogs", 302)
}

type BlogCreateController struct {
	beego.Controller
}

func (bc *BlogCreateController) Get() {
	bc.TplName = "blog-create.tpl"
}

func (bc *BlogCreateController) Post() {
	// var blog *models.Blog
	blog := new(models.Blog)
	bc.ParseForm(blog)
	fmt.Println(blog)
	o := orm.NewOrm()
	o.Insert(blog)
	bc.Redirect("/blogs", 302)
}

type BlogEditController struct {
	beego.Controller
}

type Res struct {
	Message string `json:"Message"`
}

func (bc *BlogEditController) Get() {
	id, _ := strconv.Atoi(bc.Ctx.Input.Param(":id"))
	blog := &models.Blog{Id: id}
	o := orm.NewOrm()
	err := o.Read(blog)
	if err != nil {
		fmt.Println(err)
		bc.Data["json"] = &Res{Message: "not fount"}
		bc.ServeJSON()
		return
	}
	bc.Data["blog"] = blog
	bc.TplName = "blog-edit.tpl"
}

func (bc *BlogEditController) Post() {
	id, _ := strconv.Atoi((bc.Ctx.Input.Param(":id")))
	blog := new(models.Blog)
	bc.ParseForm(blog)
	blog.Id = id
	o := orm.NewOrm()
	o.Update(blog)
	bc.Redirect("/blogs/"+string(id), 302)
}
