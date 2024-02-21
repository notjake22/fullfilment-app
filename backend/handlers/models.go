package handlers

import "main/database"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//Timestamp int64  `json:"timestamp"`
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
	Items   []database.ItemDataModel `json:"items"`
	Success bool                     `json:"success"`
	Error   string                   `json:"error"`
}

type ScanRequest struct {
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
