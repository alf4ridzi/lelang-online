package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var HubInstance = NewHub()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebsocket(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	HubInstance.Register <- conn

	go func() {
		defer func() { HubInstance.Unregister <- conn }()
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				return
			}
		}
	}()
}

func BroadcastMessage(msg string) {
	HubInstance.Broadcast <- []byte(msg)
}

func BroadcastJson(v any) {
	go func() {
		bytes, _ := json.Marshal(v)
		HubInstance.Broadcast <- bytes
	}()
}
