package smart

import (
	"errors"
	"fmt"
	"go-academy/chatroom/util"
	"net/http"

	"github.com/gorilla/websocket"
)

// NewMessageStreamHandler returns a streams endpoint handler.
func NewMessageStreamHandler(u *websocket.Upgrader) (http.HandlerFunc, error) {
	if broker == nil {
		return nil, errors.New("please run your broker with RunBroker")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := u.Upgrade(w, r, nil)
		if err != nil {
			util.LogErr("Upgrade", err)
			return
		}

		defer conn.Close()

		cli := newWebSocketClient(conn)
		broker.addClient <- cli

		defer func() {
			broker.removeClient <- cli
		}()

		util.LogInfo(fmt.Sprintf("client %s has joined chatroom", cli.ID()))
		cli.Listen()
		util.LogInfo(fmt.Sprintf("client %s has left chatroom", cli.ID()))
	}, nil
}
