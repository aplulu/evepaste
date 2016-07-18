package routers

import (
	"fareastdominions.com/evepaste/controllers"
	"github.com/astaxie/beegae"
)

func init() {
	beegae.Router("/", &controllers.PastesController{}, "*:Index;post:NewPaste")
	beegae.Router("/p/:id([0-9a-zA-Z]+)", &controllers.PastesController{}, "get:View")
	beegae.Router("/p/:id([0-9a-zA-Z]+)/raw", &controllers.PastesController{}, "get:Raw")
	beegae.Router("/p/:id([0-9a-zA-Z]+)/eft/:lang([a-z]{2})", &controllers.PastesController{}, "get:Eft")
	beegae.Router("/p/:id([0-9a-zA-Z]+)/crest", &controllers.PastesController{}, "post:Crest")
	beegae.Router("/p/:id([0-9a-zA-Z]+)/clf", &controllers.PastesController{}, "get:Clf")
	beegae.Router("/p/:id([0-9a-zA-Z]+)/dna", &controllers.PastesController{}, "get:Dna")
	beegae.Router("/p/:id([0-9a-zA-Z]+)/xml", &controllers.PastesController{}, "get:Xml")
	beegae.Router("/history", &controllers.PastesController{}, "*:History")

	beegae.Router("/accounts/login", &controllers.AccountsController{}, "*:Login")
	beegae.Router("/accounts/logout", &controllers.AccountsController{}, "*:Logout")

	beegae.Router("/legal", &controllers.PagesController{}, "*:Legal")
}
