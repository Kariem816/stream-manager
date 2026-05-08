package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Kariem816/stream-manager/internal/msgs"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func Stream(w http.ResponseWriter, r *http.Request) {
	streamIdStr := r.URL.Path[1:]
	streamId, err := strconv.Atoi(streamIdStr)
	if err != nil {
		log.Print("strconv.Atoi:", err)
		return
	}
	if streamId < 0 || streamId > 65535 {
		log.Print("invalid stream id:", streamId)
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	buf, err := msgs.MsgStreamID{StreamID: uint16(streamId)}.Buf()
	if err != nil {
		log.Print("buf:", err)
		return
	}
	defer c.Close()
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		err = c.WriteMessage(websocket.BinaryMessage, buf)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
