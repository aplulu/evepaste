package controllers

import (
	"golang.org/x/oauth2"
	"github.com/astaxie/beegae"
	"crypto/rand"
	"encoding/hex"
	"log"
	"io/ioutil"
	"encoding/json"
	"fareastdominions.com/evepaste/models/user"
	"google.golang.org/appengine"
)

var eveOAuth2 = getEVEOAuth2Config();

func getEVEOAuth2Config() *oauth2.Config {
	section := "eveoauth2"
	if appengine.IsDevAppServer() {
		section = "eveoauth2_dev"
	}
	return &oauth2.Config{
		ClientID: beegae.AppConfig.String(section + "::client_id"),
		ClientSecret: beegae.AppConfig.String(section + "::client_secret"),
		RedirectURL:  beegae.AppConfig.String(section + "::redirect_url"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://login.eveonline.com/oauth/authorize",
			TokenURL: "https://login.eveonline.com/oauth/token",
		},
		Scopes: []string{"characterFittingsWrite"},
	}
}


type AccountsController struct {
	AppController
}

func (c *AccountsController) Prepare() {
	c.AppController.Prepare()
}

func (c *AccountsController) Login() {
	if c.IsLogged() {
		c.Redirect("/", 302)
		return
	}

	e := c.Ctx.Input.Query("error")
	code := c.Ctx.Input.Query("code")
	if e != "" {
		flash := beegae.NewFlash()
		flash.Error("Unable to login")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
	} else if code != "" {
		state := c.Ctx.Input.Query("state")
		if state != c.GetSession("oauth2_state") {
			flash := beegae.NewFlash()
			flash.Error("Invalid state")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}
		token, err := eveOAuth2.Exchange(c.AppEngineCtx, code)
		if err != nil {
			log.Println(err)
			flash := beegae.NewFlash()
			flash.Error("Unable to login")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}
		log.Println(token)

		client := eveOAuth2.Client(c.AppEngineCtx, token)
		resp, err := client.Get("https://login.eveonline.com/oauth/verify")
		if err != nil {
			log.Println(err)
			flash := beegae.NewFlash()
			flash.Error("Unable to login")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			flash := beegae.NewFlash()
			flash.Error("Unable to login")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}

		user := user.User{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Println(err)
			flash := beegae.NewFlash()
			flash.Error("Unable to login")
			flash.Store(&c.Controller)
			c.Redirect("/", 302)
			return
		}

		bytes, err := json.Marshal(token)
		if err == nil {
			c.SetSession("user.token", bytes)
		}
		c.SetSessionUser(user)
		c.Redirect("/", 302)
	} else {
		b := make([]byte, 40)
		rand.Read(b)
		state := hex.EncodeToString(b)
		c.SetSession("oauth2_state", state)

		c.Redirect(eveOAuth2.AuthCodeURL(state), 302)
	}
}

func (c *AccountsController) Logout() {
	c.DelSessionUser()
	c.Redirect("/", 302)

}