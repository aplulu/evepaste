package controllers
import "github.com/beego/i18n"

type PagesController struct {
	AppController
}

func (c *PagesController) Legal() {
	c.TplName = "pages/legal.html"
	c.Data["Title"] = i18n.Tr(c.Lang, "legal")
}
