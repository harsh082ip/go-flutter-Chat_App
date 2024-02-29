package ws

import (
	"github.com/gorilla/websocket"
)

func GetwebSocketConnection(connID string) *websocket.Conn {

	conn, ok := ConnectionsMap[connID]
	if !ok {
		return nil
	}

	return conn
}

func AddWebSocketConnection(connID string, conn *websocket.Conn) {
	ConnectionsMap[connID] = conn
}

func RemoveWebSocketConnection(connID string) {
	delete(ConnectionsMap, connID)
}
