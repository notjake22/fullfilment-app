package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/database"
	"main/handlers"
	"main/msHandlers"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	// supported builds
	//b1 := router.Group("/")
	b2 := router.Group("/b2")
	b3 := router.Group("/b3")

	r := newGroupGroup([]*gin.RouterGroup{b2, b3})

	database.InitDbConnection()
	defer database.DBConn.Close()

	// init
	r.handle(http.MethodPost, "/login", handlers.Login)
	r.handle(http.MethodGet, "/build/check-version", handlers.CheckBuildNumber)

	r.handle(http.MethodPost, "/add-new-client", handlers.SetNewClient)
	r.handle(http.MethodPost, "/edit-client", handlers.EditClient)
	r.handle(http.MethodGet, "/user-list", handlers.GetUsers)

	r.handle(http.MethodGet, "/inventory/get-all", handlers.GetAllInventoryItems)
	r.handle(http.MethodDelete, "/inventory/delete-item", handlers.DeleteInventoryItem)

	r.handle(http.MethodPost, "/scan/in", handlers.ScanIn)
	r.handle(http.MethodPost, "/scan/out", handlers.ScanOut)
	r.handle(http.MethodPost, "/scan/out/cancel", handlers.ScanOutCancel)
	r.handle(http.MethodPost, "/scan/in/cancel", handlers.ScanInCancel)
	r.handle(http.MethodPost, "/scan/new-item", handlers.AddCustomItem)

	r.handle(http.MethodGet, "/service-settings", handlers.GetServSettings)
	r.handle(http.MethodPost, "/service-settings/set-new", handlers.PutServSettings)

	r.handle(http.MethodGet, "/pricing/all", msHandlers.GetAllPricing)
	r.handle(http.MethodGet, "/pricing/id/:Id", msHandlers.GetPrice)
	r.handle(http.MethodPost, "/pricing/new", msHandlers.NewPrice)
	r.handle(http.MethodPost, "/pricing/update", msHandlers.UpdatePrice)

	log.Println("Starting HTTP endpoints")
	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
		return
	}
}

type GroupGroup struct {
	groups []*gin.RouterGroup
}

func newGroupGroup(groups []*gin.RouterGroup) GroupGroup {
	return GroupGroup{
		groups,
	}
}

func (g *GroupGroup) handle(method string, path string, handler gin.HandlerFunc) {
	for _, group := range g.groups {
		group.Handle(method, path, handler)
	}
}
