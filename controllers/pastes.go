package controllers

import (
	"log"
	"strings"
	"github.com/astaxie/beegae"
	"fareastdominions.com/evepaste/utils"
	"fareastdominions.com/evepaste/eve/parser"
	"github.com/beego/i18n"
	"fareastdominions.com/evepaste/models/paste"
	"fareastdominions.com/evepaste/eve/entity"
	"fareastdominions.com/evepaste/eve/fit"
	"fareastdominions.com/evepaste/models"
)

type PastesController struct {
	AppController
}

func (c *PastesController) Prepare() {
	c.AppController.Prepare()
	c.Data["Systems"] = strings.Split(beegae.AppConfig.String("eve::systems"), "|")
}

/**
 * Home
 */
func (c *PastesController) Index() {
	c.TplName = "pastes/index.html"
	c.Data["IsHome"] = true
}

/**
 * View paste
 */
func (c *PastesController) View() {
	c.TplName = "pastes/view.html"

	id := utils.DecodeBase62(c.Ctx.Input.Param(":id"))
	p, err := paste.GetPaste(c.AppEngineCtx, id)
	if err != nil {
		log.Println(err)
		c.Abort("404")
		return
	}

	p.Items = parser.ApplyI18nName(p.Items, c.Lang)
	if p.Type == "dscan" {
		c.Data["IsDscan"] = true
		c.Data["ScanResult"] = parser.ScanItems(p.Items, p.ScanSystemID, c.Lang)
	} else if p.Type == "fit" {
		c.Data["IsFit"] = true
		c.Data["Fit"] = entity.NewFitFromItems(p.Items)
	} else {
		groups := p.GetGroupedItems()
		if len(groups) > 0 {
			c.Data["IsGrouped"] = true
			c.Data["Groups"] = groups

		}
	}

	c.Data["Title"] = p.Type + ":" + p.EncodedID
	c.Data["Paste"] = p
}

/**
 * Raw paste
 */
func (c *PastesController) Raw() {
	id := utils.DecodeBase62(c.Ctx.Input.Param(":id"))
	p, err := paste.GetPaste(c.AppEngineCtx, id)
	if err != nil {
		log.Println(err)
		c.Abort("404")
		return
	}

	c.Ctx.Output.Header("Content-Type", "text/plain; charset=utf-8")
	c.Ctx.WriteString(p.Text)
}

/**
 * Export eft
 */
func (c *PastesController) Eft() {
	id := utils.DecodeBase62(c.Ctx.Input.Param(":id"))
	lang := c.Ctx.Input.Param(":lang")

	p, err := paste.GetPaste(c.AppEngineCtx, id)
	if err != nil || p.Type != "fit" {
		log.Println(err)
		c.Abort("404")
		return
	}

	f := entity.NewFitFromItems(p.Items)

	c.Ctx.Output.Header("Content-Type", "text/plain; charset=utf-8")
	c.Ctx.WriteString(fit.ExportEFT(f, lang))
}

/**
 * Export clf
 */
func (c *PastesController) Clf() {
	id := utils.DecodeBase62(c.Ctx.Input.Param(":id"))

	p, err := paste.GetPaste(c.AppEngineCtx, id)
	if err != nil || p.Type != "fit" {
		log.Println(err)
		c.Abort("404")
		return
	}

	f := entity.NewFitFromItems(p.Items)

	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.WriteString(fit.ExportCLF(f))
}

/**
 * Export dna
 */
func (c *PastesController) Dna() {
	id := utils.DecodeBase62(c.Ctx.Input.Param(":id"))

	p, err := paste.GetPaste(c.AppEngineCtx, id)
	if err != nil || p.Type != "fit" {
		log.Println(err)
		c.Abort("404")
		return
	}

	f := entity.NewFitFromItems(p.Items)

	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.WriteString(fit.ExportDNA(f))
}

/**
 * Export xml
 */
func (c *PastesController) Xml() {

}

/**
 * Export crest
 */
func (c *PastesController) Crest() {
	id := utils.DecodeBase62(c.Ctx.Input.Param(":id"))

	p, err := paste.GetPaste(c.AppEngineCtx, id)
	if err != nil || p.Type != "fit" {
		log.Println(err)
		c.Abort("404")
		return
	}

	if !c.IsLogged() {
		c.Ctx.Output.SetStatus(401)
		c.respondJSON(models.Error{
			Error: "login_required",
			ErrorDescription: i18n.Tr(c.Lang, "login_required"),
		})
		return
	}

	token := c.GetSessionToken()
	if token == nil {
		c.Ctx.Output.SetStatus(401)
		c.respondJSON(models.Error{
			Error: "login_required",
			ErrorDescription: i18n.Tr(c.Lang, "login_required"),
		})
		return
	}

	client := getEVEOAuth2Config().Client(c.AppEngineCtx, token)

	f := entity.NewFitFromItems(p.Items)
	payload := fit.ExportCREST(f, p.EncodedID)
	err = fit.PostCREST(client, c.User, payload)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.respondJSON(models.Error{
			Error: "error",
			ErrorDescription: err.Error(),
		})
	} else {
		c.respondJSON(models.Message{
			Message: i18n.Tr(c.Lang, "crest_fitting_exported"),
		})
	}

}

/**
 * Submit new paste
 */
func (c *PastesController) NewPaste() {
	flash := beegae.NewFlash()

	var text string
	var systemId int
	var compressOre string
	var decompressOre string
	c.Ctx.Input.Bind(&text, "text")
	c.Ctx.Input.Bind(&systemId, "system_id")
	c.Ctx.Input.Bind(&compressOre, "compress_ore")
	c.Ctx.Input.Bind(&decompressOre, "decompress_ore")

	p := paste.NewPaste(text, systemId)
	if !p.IsValid() {
		flash.Error(i18n.Tr(c.Lang, "unable_parse"))
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/")
		return
	}

	if compressOre != "" {
		p.Items = parser.CompressItems(p.Items)
	} else if decompressOre != "" {
		p.Items = parser.DecompressItems(p.Items)
	}

	if p.Type == "dscan" {
		p.ScanSystemID = int(parser.ScanSolarSystem(p.Items))
		if p.ScanSystemID > 0 {
			log.Printf("Solar System ID: %d", p.ScanSystemID)
		}
	}

	p.IPAddr = c.GetRemoteAddr()
	if c.IsLogged() {
		p.User = c.User
	}

	p.FetchMarketPrice(c.AppEngineCtx)

	p, err := p.Save(c.AppEngineCtx)
	if err != nil {
		log.Println(err)
		flash.Error(i18n.Tr(c.Lang, "unable_save_paste"))
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/")
		return
	}

	//c.Ctx.Redirect(302, "/p/" + strconv.FormatInt(p.ID, 10))
	c.Ctx.Redirect(302, "/p/" + utils.EncodeBase62(p.ID))
}

/**
 * Paste history
 */
func (c *PastesController) History() {
	if !c.ValidateLogged() {
		return
	}
	c.TplName = "pastes/history.html"
	c.Data["IsHistory"] = true

	pastes, err := paste.GetPasteHistory(c.AppEngineCtx, c.User)
	if err != nil {
		log.Println(err)
		flash := beegae.NewFlash()
		flash.Error("Unable to retrive history")
		return
	}

	c.Data["Pastes"] = pastes
	c.Data["History"] = i18n.Tr(c.Lang, "history")
}
