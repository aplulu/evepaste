package main

import (
	"github.com/astaxie/beegae"
	_ "github.com/astaxie/beegae/session/appengine"
	_ "fareastdominions.com/evepaste/routers"
	"fareastdominions.com/evepaste/controllers"
	"fareastdominions.com/evepaste/template"
	"fareastdominions.com/evepaste/eve/repository"
)

func init() {
	repository.LoadTables("data/")
	template.InitTemplate()
	controllers.InitController()
	beegae.Run()
}
