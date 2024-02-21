package database

const (
	baseApiUrl string = "https://glum-backend-production.up.railway.app/b2"
)

type GeneralResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
}

type UsersResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Users   []User `json:"users"`
}

type User struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Phone                string `json:"phone"`
	Balance              int    `json:"balance"`
	CurrentlyStoredItems int    `json:"currentlyStoredItems"`
	TotalOrderHistory    int    `json:"totalOrderHistory"`
}

type InventoryResponse struct {
	Items   []ItemDataModel `json:"items"`
	Success bool            `json:"success"`
	Error   string          `json:"error"`
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

type ScanInRequest struct {
	UPC      string `json:"upc"`
	ClientID string `json:"clientID"`
}

type ScanInResponse struct {
	ID       string `json:"ID"`
	ItemName string `json:"itemName"`
	ImageUri string `json:"imageUri"`
}

type ScanInCancelRequest struct {
	ItemID string `json:"itemID"`
}

type ServiceSettings struct {
	ItemPrice int `json:"itemPrice"`
}

type CustomUPCModel struct {
	ItemName string `json:"itemName"`
	ImageUri string `json:"imageUri"`
	UPC      string `json:"upc"`
}
