package paste

import (
	"time"
	"github.com/mjibson/goon"
	"google.golang.org/appengine/datastore"
	"fareastdominions.com/evepaste/models/user"
	"fareastdominions.com/evepaste/eve/parser"
	"golang.org/x/net/context"
)


/**
 * New paste from text
 */
func NewPaste(text string, systemId int) *Paste {
	paste := &Paste{
		Text: text,
		MarketSystemID: systemId,
		Created: time.Now(),
	}
	paste.Type, paste.Items = parser.Parse(text)

	return paste
}

/**
 * Read paste from datastore
 */
func GetPaste(c context.Context, id int64) (*Paste, error) {
	paste := Paste{
		ID: id,
	}
	g := goon.FromContext(c)
	err := g.Get(&paste)
	if err != nil {
		return nil, err
	}

	err = paste.unserialize()
	if err != nil {
		return nil, err
	}

	return &paste, nil
}

/**
 * Read paste histories
 */
func GetPasteHistory(c context.Context, user user.User) ([]Paste, error) {
	q := datastore.NewQuery("Paste").Filter("User.ID =", user.ID).Order("-Created").Limit(100)
	g := goon.FromContext(c)

	var pastes []Paste
	_, err := g.GetAll(q, &pastes)
	if err == nil {
		for i, _ := range pastes {
			pastes[i].unserialize()
		}
	}

	return pastes, err
}
