package upc

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func (u *LookUp) findItemInfo(doc *html.Node) error {
	itemNameNode := htmlquery.FindOne(doc, `/html/body/div/section[2]/div/div/div/div[2]/div/div[1]/h2`)
	if itemNameNode == nil {
		return errors.New("error finding title node")
	} else {
		u.item.ItemName = htmlquery.InnerText(itemNameNode)
	}

	imageUriNode := htmlquery.FindOne(doc, `/html/body/div/section[2]/div/div/div/div[2]/div/div[2]/div[1]/div/div[1]/div[1]/img`)
	if imageUriNode == nil {
		return errors.New("error finding uri node")
	} else {
		u.item.ItemImageUri = imageUriNode.Attr[1].Val
	}
	return nil
}
