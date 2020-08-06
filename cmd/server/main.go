package main

import (
	"github.com/astaxie/beegae"
	_ "github.com/astaxie/beegae/session/appengine"
	_ "github.com/evepaste/evepaste/pkg/routers"
	"github.com/evepaste/evepaste/pkg/controllers"
	"github.com/evepaste/evepaste/pkg/template"
	"github.com/evepaste/evepaste/pkg/eve/repository"
	"google.golang.org/appengine"
)

func main() {
	if !appengine.IsDevAppServer() {
		beegae.BConfig.RunMode = "prod"
	}
	repository.LoadTables("data/")
	template.InitTemplate()
	controllers.InitController()
	beegae.Run()
}
