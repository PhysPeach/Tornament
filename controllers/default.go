package controllers

import (
	"github.com/astaxie/beego"
)

// RootController はトップページのコントローラ
type RootController struct {
	beego.Controller
}

// Get はトップページを表示する
//
// GET /
func (c *RootController) Get() {
	// /views/index.tpl を表示する
	c.TplName = "index.tpl"
}
