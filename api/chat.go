package api

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	Room    string `json:"room"`
	Author  string `json:"author"`
	Content string `json:"content"`
	IsMe    bool   `json:"isMe"`
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	rooms = make(map[string]map[*websocket.Conn]bool)

	roomMux sync.Mutex
)

func (s *Server) HandlerMessage(c *gin.Context) {
	room := c.Param("room")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	roomMux.Lock()

	if _, ok := rooms[room]; !ok {
		rooms[room] = make(map[*websocket.Conn]bool)
	}

	rooms[room][conn] = true

	roomMux.Unlock()

	go readMessage(room, conn)
}

func readMessage(room string, conn *websocket.Conn) {
	defer func() {
		conn.Close()
		roomMux.Lock()
		delete(rooms[room], conn)

		roomMux.Unlock()
	}()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)

		if err != nil {
			fmt.Println(err)
			break
		}

		go broadcastMessage(room, msg)
	}
}

func broadcastMessage(room string, msg Message) {
	roomMux.Lock()

	for conn := range rooms[room] {
		err := conn.WriteJSON(msg)

		if err != nil {
			fmt.Println(err)
			return
		}

	}

	roomMux.Unlock()
}
