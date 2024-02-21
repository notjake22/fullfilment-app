package database

import "log"

func GetItemByID(id string) (ItemDataModel, error) {
	row := DBConn.QueryRow(`SELECT "item_name", "item_image_uri", "upc_str", "owner_name", "owner_id" 
								FROM "inventory" WHERE id = $1`, id)
	if row.Err() != nil {
		return ItemDataModel{}, row.Err()
	}

	var itemData ItemDataModel
	err := row.Scan(&itemData.ItemName, &itemData.ImageUri, &itemData.UPC, &itemData.OwnerName, &itemData.OwnerID)
	if err != nil {
		return ItemDataModel{}, err
	}
	itemData.ID = id

	return itemData, nil
}

func GetAllInventoryItems() ([]ItemDataModel, error) {
	rows, err := DBConn.Query(`SELECT "id", "item_name", "item_image_uri", "upc_str", "owner_name", "owner_id", "currently_stored" FROM "inventory"`)
	if err != nil {
		return nil, err
	}

	var items []ItemDataModel
	var item ItemDataModel
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.ItemName, &item.ImageUri, &item.UPC, &item.OwnerName, &item.OwnerID, &item.CurrentlyStored)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func CheckUPCByOwner(userId, upc string) (ItemDataModel, error) {
	rows, err := DBConn.Query(`SELECT "id","item_name", "item_image_uri", "owner_name"
									FROM "inventory" 
									WHERE "owner_id" = $1 AND "upc_str" = $2 AND "currently_stored" = $3`, userId, upc, true)
	if err != nil {
		log.Println("Query issue: ", err)
		return ItemDataModel{}, err
	}

	var item ItemDataModel
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.ItemName, &item.ImageUri, &item.OwnerName)
		if err != nil {
			log.Println("Scan issue: ", err)
			return ItemDataModel{}, err
		}
		break
	}

	item.UPC = upc
	item.OwnerID = userId
	item.CurrentlyStored = true

	return item, nil
}

func GetUserItems(userId string) ([]ItemDataModel, error) {
	rows, err := DBConn.Query(`SELECT "id", "item_name", "item_image_uri", "upc_str", "owner_name", "owner_id", "currently_stored"
									FROM "inventory" WHERE "owner_id" = $1`, userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var items []ItemDataModel
	var item ItemDataModel
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.ItemName, &item.ImageUri, &item.UPC, &item.OwnerName, &item.OwnerID, &item.CurrentlyStored)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func SetNewItem(data ItemDataModel) (string, error) {
	var itemID string
	err := DBConn.QueryRow(`INSERT INTO "inventory" 
    ("item_name", "item_image_uri", "upc_str", "owner_name", "owner_id", "currently_stored", upc)
    values($1, $2, $3, $4, $5, $6, $7)
    RETURNING id`,
		data.ItemName, data.ImageUri, data.UPC, data.OwnerName, data.OwnerID, data.CurrentlyStored, 0,
	).Scan(&itemID)
	if err != nil {
		log.Println("Error storing new inventory item: ", err)
		return "", err
	}

	return itemID, nil
}

func DeleteItem(id string) error {
	_, err := DBConn.Exec(`DELETE FROM "inventory" WHERE id = $1`, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func MarkItemOut(id string, mark bool) error {
	_, err := DBConn.Exec(`UPDATE "inventory" SET "currently_stored" = $1 WHERE id = $2`, mark, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
