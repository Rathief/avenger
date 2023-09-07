package main

import (
	"avenger/config"
	"avenger/handler"
	"avenger/repo"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	ah := handler.AuthHandler{
		DBHandler: repo.DBHandler{DB: db},
	}
	e := echo.New()
	e.POST("/register", ah.Register)
	e.POST("/login", ah.Login)

	e.GET("/stores", ah.GetStores)
	e.GET("/stores/:id", ah.GetStoreByID)

	e.GET("/ws", hello)

	e.Logger.Fatal(e.Start(":8080"))
}
