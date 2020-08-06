package fit

import (
	"github.com/evepaste/evepaste/pkg/eve/entity"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/evepaste/evepaste/pkg/models/user"
	"strings"
	"io/ioutil"
"log"
	"time"
	"errors"
)

type CRESTFit struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Items []CRESTFitItem `json:"items"`
	Ship CRESTType `json:"ship"`
}

type CRESTFitItem struct {
	Type CRESTType `json:"type"`
	Flag int `json:"flag"`
	Quantity int `json:"quantity"`
}

type CRESTType struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

var crestSlotMap = map[string]int {
	"low": 11,
	"med": 19,
	"high": 27,
	"rig": 92,
	"subsystem": 125,
	"drone": 87,
	"other": 5,
}

func newCRESTType32(id int32) CRESTType {
	return CRESTType{
		Id: int(id),
		Href: fmt.Sprintf("https://crest-tq.eveonline.com/inventory/types/%d/", id),
	}
}

func ExportCREST(f *entity.Fit, pasteId string) string {
	items := make([]CRESTFitItem, 0)
	for _, slot := range f.Slots {
		if slot.Name == "ship" {
			continue
		}

		baseSlot := crestSlotMap[slot.Name]

		for i, item := range slot.Items {
			if slot.Name == "drone" || slot.Name == "other" {
				items = append(items, CRESTFitItem{
					Type: newCRESTType32(item.TypeId),
					Quantity: item.Quantity,
					Flag: baseSlot,
				})
			} else {
				if i > 7 {
					break
				}

				items = append(items, CRESTFitItem{
					Type: newCRESTType32(item.TypeId),
					Quantity: item.Quantity,
					Flag: baseSlot + i,
				})
			}
		}
	}


	crestFit := &CRESTFit{
		Name: "EVEPaste import: " + f.Ship.Name,
		Description: "https://eve-paste.appspot.com/p/" + pasteId,
		Ship: newCRESTType32(f.Ship.TypeId),
		Items: items,
	}

	bytes, err := json.Marshal(crestFit)
	if err != nil {
		return ""
	} else {
		return string(bytes)
	}
}

func PostCREST(c *http.Client, u user.User, payload string) error {
	c.Timeout = 10 * time.Second
	req, err := http.NewRequest("POST", fmt.Sprintf("https://crest-tq.eveonline.com/characters/%d/fittings/", u.ID), strings.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		log.Println(payload)
		log.Println(resp.StatusCode)
		log.Println(string(body))

		m := entity.CrestMessage{}
		err = json.Unmarshal(body, &m)
		if m.Message != "" {
			return errors.New(m.Message)
		} else {
			return errors.New("Failed to write fitting.")
		}
	}

	return nil
}