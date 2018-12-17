package naive

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

		reg := registration{conn: conn, id: util.RandStringID(15)}
		broker.addConn <- reg

		defer func() {
			broker.removeConn <- reg
		}()

		util.LogInfo(fmt.Sprintf("client:%s has joined chatroom", reg.id))

		for {
			p := Payload{}
			err := conn.ReadJSON(&p)
			if err != nil &&
				websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				util.LogInfo(fmt.Sprintf("client %s has left the chatroom", reg.id))
				return
			}

			if err != nil {
				// This is an abnormal error, important to log it for debugging.
				util.LogErr("ReadJSON", err)
				return
			}

			broker.broadcast <- p
		}
	}, nil
}
