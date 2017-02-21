package main

import (
	"github.com/astaxie/beegae"
	_ "github.com/astaxie/beegae/session/appengine"
	_ "fareastdominions.com/evepaste/routers"
	"fareastdominions.com/evepaste/controllers"
	"fareastdominions.com/evepaste/template"
	"fareastdominions.com/evepaste/eve/repository"
	"google.golang.org/appengine"
)

func init() {
	if !appengine.IsDevAppServer() {
		beegae.BConfig.RunMode = "prod"
	}
	repository.LoadTables("data/")
	template.InitTemplate()
	controllers.InitController()
	beegae.Run()
}
