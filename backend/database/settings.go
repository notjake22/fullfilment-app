package database

import "log"

func QueryAdminSettings() (AdminSettings, error) {
	row := DBConn.QueryRow(`SELECT "item_price" FROM "admin-settings-0" WHERE id = $1`, 1)
	if row.Err() != nil {
		log.Println(row.Err())
		return AdminSettings{}, row.Err()
	}

	var settings AdminSettings
	err := row.Scan(&settings.ItemPrice)
	if err != nil {
		log.Println(err)
		return AdminSettings{}, err
	}

	return settings, nil
}

func SetNewAdminSettings(settings AdminSettings) error {
	_, err := DBConn.Exec(`UPDATE "admin-settings-0" SET "item_price" = $1 WHERE id = $2`, settings.ItemPrice, 1)
	if err != nil {
		return err
	}

	return nil
}
