package database

import (
	"log"
)

func GetCustomItem(upc string) (CustomUPCModel, error) {
	row := DBConn.QueryRow(`SELECT "id", "name","image_uri" FROM "custom-items" WHERE upc = $1`, upc)
	if row.Err() != nil {
		return CustomUPCModel{}, row.Err()
	}

	var item CustomUPCModel
	err := row.Scan(&item.ID, &item.ItemName, &item.ImageUri)
	if err != nil {
		return CustomUPCModel{}, err
	}

	item.UPC = upc
	return item, nil
}

func SetNewCustomItem(itemName, itemUri, upc string) error {
	_, err := DBConn.Exec(`insert into "custom-items" ("name", "image_uri", "upc") values($1, $2, $3)`, itemName, itemUri, upc)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
