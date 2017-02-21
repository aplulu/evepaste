package controllers

import (
	"github.com/astaxie/beegae"
	"github.com/beego/i18n"
	"strings"
	"fareastdominions.com/evepaste/models/user"
	"time"
	"golang.org/x/oauth2"
	"encoding/json"
	"google.golang.org/appengine"
)

type AppController struct {
	beegae.Controller
	i18n.Locale
	User user.User
}

func (c *AppController) Prepare() {
	c.Data["StartTime"] = time.Now()

	isSecure := c.Ctx.Request.TLS != nil || c.Ctx.Request.Header.Get("X-Forwarded-Proto") == "https";
	if !isSecure && !appengine.IsDevAppServer() {
		url := "https://" + c.Ctx.Request.Host + c.Ctx.Request.RequestURI

		c.Redirect(url, 301)
		return

	}

	c.Controller.Prepare()
	beegae.ReadFromRequest(&c.Controller)

	c.User = c.GetSessionUser()
	c.Layout = "default.html"
	c.Data["Ctx"] = c.Ctx
	c.Data["Gctx"] = c.AppEngineCtx
	c.Data["User"] = c.User
	c.Data["Logged"] = c.IsLogged()
	c.Data["StaticUrl"] = beegae.AppConfig.String("staticUrl")

	if c.setLangVer() {
		i := strings.Index(c.Ctx.Request.RequestURI, "?")
		c.Redirect(c.Ctx.Request.RequestURI[:i], 302)
		return
	}
}

func (c *AppController) GetRemoteAddr() string {
	ips := c.Ctx.Input.Proxy()
	if len(ips) > 0 && ips[0] != "" {
		rip := strings.Split(ips[0], ":")
		if len(rip) < 3 {
			return rip[0]
		} else {
			return ips[0]
		}
	}

	ip := c.Ctx.Input.Context.Request.RemoteAddr
	ips = strings.Split(ip, ":")
	if strings.HasPrefix(ip, "[") {
		return ip[1:strings.LastIndex(ip, ":") - 1]
	} else if len(ips) == 2 {
		return ips[0]
	} else {
		return ip
	}
}

func (c *AppController) IsLogged() bool {
	return c.User.ID > 0
}

func (c *AppController) ValidateLogged() bool {
	if !c.IsLogged() {
		flash := beegae.NewFlash()
		flash.Error(i18n.Tr(c.Lang, "unauthorized"))
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/")
		return false
	}
	return true
}

func (c *AppController) SetSessionUser(u user.User) {
	c.SetSession("user.id", u.ID)
	c.SetSession("user.name", u.Name)
	c.SetSession("user.owner_hash", u.OwnerHash)
}

func (c *AppController) GetSessionUser() user.User {
	id, _ := c.GetSession("user.id").(int)
	name, _ := c.GetSession("user.name").(string)
	ownerHash, _ := c.GetSession("user.owner_hash").(string)
	return user.User{
		ID: id,
		Name: name,
		OwnerHash: ownerHash,
	}
}

func (c *AppController) GetSessionToken() *oauth2.Token {
	bytes, ok := c.GetSession("user.token").([]byte)
	if ok {
		token := oauth2.Token{}
		err := json.Unmarshal(bytes, &token)
		if err == nil {
			return &token
		}
	}
	return nil
}

func (c *AppController) DelSessionUser() {
	c.DelSession("user.id")
	c.DelSession("user.name")
	c.DelSession("user.owner_hash")
	c.DelSession("user.token")
}

func (c *AppController) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := c.Input().Get("lang")
	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = c.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		acceptLanguages := strings.SplitN(c.Ctx.Request.Header.Get("Accept-Language"), ",", -1)
		for _, l := range acceptLanguages {
			l = strings.SplitN(l, ";", 2)[0]
			if i18n.IsExist(l) {
				lang = l
				break
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "en"
		isNeedRedir = false
	}

	curLang := langType{
		Lang: lang,
	}

	// Save language information in cookies.
	if !hasCookie {
		c.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
	}

	restLangs := make([]*langType, 0, len(langTypes)-1)
	for _, v := range langTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}

	// Set language properties.
	c.Lang = lang
	c.Data["Lang"] = curLang.Lang
	c.Data["CurLang"] = curLang.Name
	c.Data["RestLangs"] = restLangs

	return isNeedRedir
}

func (c *AppController) respondJSON(v interface{}) {
	bytes, err := json.Marshal(v)
	if err != nil {
		c.Abort("500")
		return
	}

	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.Output.Body(bytes)
}