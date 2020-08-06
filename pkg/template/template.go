package template

import (
	"github.com/astaxie/beegae"
	"github.com/beego/i18n"
	"html/template"
	"strings"
	"strconv"
	"fmt"
	"github.com/evepaste/evepaste/pkg/eve/entity"
	"time"
	"github.com/xeonx/timeago"
	"github.com/evepaste/evepaste/pkg/models/paste"
	"golang.org/x/net/context"
	"encoding/json"
	"runtime"
)

func InitTemplate() {
	beegae.AddFuncMap("css", css)
	beegae.AddFuncMap("javascript", javascript)
	beegae.AddFuncMap("comma", comma)
	beegae.AddFuncMap("commaF", commaF)
	beegae.AddFuncMap("aboutPrice", aboutPrice)
	beegae.AddFuncMap("itemPrice", itemPrice)
	beegae.AddFuncMap("pasteTotal", pasteTotal)
	beegae.AddFuncMap("systemName", systemName)
	beegae.AddFuncMap("systemNameI", systemNameI)
	beegae.AddFuncMap("time", formatTime)
	beegae.AddFuncMap("processTime", processTime)
	beegae.AddFuncMap("goVersion", goVersion)
	beegae.AddFuncMap("appengineInfo", appengineInfo)
	beegae.AddFuncMap("encodeJSON", encodeJSON)
	beegae.AddFuncMap("jsLang", jsLang)
	beegae.AddFuncMap("eveLangCodes", eveLangCodes)

	beegae.AddFuncMap("i18n", i18n.Tr)
	beegae.AddFuncMap("i18nH", i18nH)
}

func css(path string) template.HTML {
	h := "<link rel=\"stylesheet\" type=\"text/css\" href=\"" + beegae.AppConfig.String("staticUrl") + "/css/" + path + "\"/>"
	return template.HTML(h)
}

func javascript(path string) template.HTML {
	h := "<script type=\"text/javascript\" src=\"" + beegae.AppConfig.String("staticUrl") + "/js/" + path + "\"></script>"
	return template.HTML(h)
}

func comma(v int) string {
	sign := ""
	if v < 0 {
		sign = "-"
		v = 0 - v
	}
	v2 := int64(v)

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v2 > 999 {
		parts[j] = strconv.FormatInt(v2%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v2 = v2 / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v2))
	return sign + strings.Join(parts[j:], ",")
}

func commaF(f float64) string {
	out := fmt.Sprintf("%.2f", f)
	parts := strings.SplitN(out, ".", 2)
	v, err := strconv.Atoi(parts[0])
	if err != nil {
		return out
	} else {
		return comma(v) + "." + parts[1]
	}
}

func aboutPrice(lang string, f float64) string {
	if f > 1000000000000 {
		return commaF(f / 1000000000000) + " Trillion"
	} else if f > 1000000000 {
		return commaF(f / 1000000000) + " Billion"
	} else if f > 1000000 {
		return commaF(f / 1000000) + " Million"
	} else {
		return commaF(f)
	}
}

func itemPrice(item entity.Item, mode string) template.HTML {
	var sell float64
	var buy float64

	if mode == "total" {
		sell = item.Prices.Sell.Price * float64(item.Quantity)
		buy = item.Prices.Buy.Price * float64(item.Quantity)
	} else if mode == "volume" {
		sell = item.Prices.Sell.Price * (1 / item.Volume)
		buy = item.Prices.Buy.Price * (1 / item.Volume)
	} else {
		sell = item.Prices.Sell.Price
		buy = item.Prices.Buy.Price

	}

	h := commaF(sell) + "<br/>" + commaF(buy)
	return template.HTML(h)
}

func pasteTotal(p paste.Paste) template.HTML {
	h := commaF(p.TotalSellPrice) + "<br/>" + commaF(p.TotalBuyPrice)
	return template.HTML(h)
}

func systemName(lang string, systemId string) string {
	return i18n.Tr(lang, "system." + systemId)
}

func systemNameI(lang string, systemId int) string {
	return i18n.Tr(lang, "system." + strconv.Itoa(systemId))
}

func formatTime(lang string, t time.Time) string {
	var config timeago.Config

	if lang == "ja" {
		config = timeago.Config{
			PastPrefix:   "",
			PastSuffix:   "前",
			FuturePrefix: "",
			FutureSuffix: "",

			Periods: []timeago.FormatPeriod{
				timeago.FormatPeriod{time.Second, "数秒", "%d秒"},
				timeago.FormatPeriod{time.Minute, "1分", "%d分"},
				timeago.FormatPeriod{time.Hour, "1時間", "%d時間"},
				timeago.FormatPeriod{timeago.Day, "1日", "%d日"},
				timeago.FormatPeriod{timeago.Month, "1ヶ月", "%dヶ月"},
				timeago.FormatPeriod{timeago.Year, "1年", "%d年"},
			},

			Zero: "数秒",

			Max:           73 * time.Hour,
			DefaultLayout: "2006/01/02",
		}
	} else {
		config = timeago.English
	}

	return config.Format(t)
}

func processTime(start time.Time) string {
	end := time.Now()
	return strconv.Itoa(int(end.Sub(start).Nanoseconds() / int64(time.Millisecond))) + "ms"
}

func goVersion() string {
	return runtime.Version()
}

func i18nH(lang, format string, args ...interface{}) template.HTML {
	return template.HTML(i18n.Tr(lang, format, args))
}

func appengineInfo(ctx context.Context) string {
	return ""
}

func encodeJSON(v interface{}) template.JS {
	out := "{}"
	b, err := json.Marshal(v)
	if err == nil {
		out = string(b)
	}
	return template.JS(out)
}

func jsLang(lang string, keys string) template.JS {
	ks := strings.SplitN(keys, ",", -1)

	langs := make(map[string]string)

	for _, k := range ks {
		langs[k] = i18n.Tr(lang, k)
	}

	out := "{}"
	b, err := json.Marshal(langs)
	if err == nil {
		out = string(b)
	}
	return template.JS(out)
}

func eveLangCodes() string {
	return beegae.AppConfig.String("lang::evetypes")
}