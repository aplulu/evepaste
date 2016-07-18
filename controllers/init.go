package controllers

import (
	"strings"
	"github.com/astaxie/beegae"
	"github.com/beego/i18n"
	"log"
)

var langTypes []*langType
type langType struct {
	Lang, Name string
}

func InitController() {
	langs := strings.Split(beegae.AppConfig.String("lang::types"), "|")
	names := strings.Split(beegae.AppConfig.String("lang::names"), "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		if err := i18n.SetMessage(lang, "locale/" + lang + ".ini"); err != nil {
			log.Println("Fail to set message file: " + err.Error())
			return
		}
	}
}