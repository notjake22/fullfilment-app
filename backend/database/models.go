package database

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Balance int    `json:"balance"`
}

type AdminSettings struct {
	ItemPrice int `json:"itemPrice"` // first 2 digits as cent values
}

type ItemDataModel struct {
	ID              string `json:"id"`
	ItemName        string `json:"itemName"`
	ImageUri        string `json:"imageUri"`
	UPC             string `json:"upc"`
	OwnerName       string `json:"ownerName"`
	OwnerID         string `json:"ownerID"`
	CurrentlyStored bool   `json:"currentlyStored"`
}

type CustomUPCModel struct {
	ID       string `json:"id"`
	ItemName string `json:"itemName"`
	ImageUri string `json:"imageUri"`
	UPC      string `json:"upc"`
}
