package main

import (
	"changeme/database"
	"changeme/load"
	"context"
	"log"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CheckLogin(user string, pass string) bool {
	login, err := load.Login(user, pass)
	if err != nil {
		log.Println(err)
		return false
	}
	return login
}

func (a *App) CheckBuildNumber() int {
	return load.CheckBuild()
}

func (a *App) GetUserTableInfo() ([]database.User, error) {
	return database.GetAllUsers()
}

func (a *App) SetNewUser(user database.User) bool {
	return database.SetNewClient(user)
}

func (a *App) EditClient(user database.User) bool {
	return database.EditClient(user)
}

func (a *App) GetInventoryItems() ([]database.ItemDataModel, error) {
	return database.GetAllItems()
}

func (a *App) DeleteInventoryItem(itemId string) error {
	return database.DeleteIvenItem(itemId)
}

func (a *App) ScanInHandler(data database.ScanInRequest) (database.ScanInResponse, error) {
	return database.ScanInItem(data)
}

func (a *App) ScanCancel(id string) error {
	return database.CancelScanItem(id)
}

func (a *App) ScanOutHandler(data database.ScanInRequest) (database.ScanInResponse, error) {
	return database.ScanOutItem(data)
}

func (a *App) ScanOutCancel(id string) error {
	return database.CancelScanOutItem(id)
}

func (a *App) AddNewCustomItem(item database.CustomUPCModel) error {
	return database.SetNewCustomItem(item)
}

func (a *App) GetServiceSettings() (database.ServiceSettings, error) {
	return database.GetServiceSettings(a.ctx)
}

func (a *App) SetNewServiceSettings(settings database.ServiceSettings) error {
	return database.SetNewServiceSettings(settings, a.ctx)
}
