package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
)

type WSClient struct {
	conn *websocket.Conn
	name string
}

func (wc *WSClient) ConnectWSClient(host string, botId int) (err error) {
	u := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   "/",
	}
	wc.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	wc.generateName(botId)
	fmt.Printf("WS-client connected! name: %s url: %s \n", wc.name, u.String())
	return nil
}

func (wc *WSClient) generateName(id int) {
	wc.name = fmt.Sprintf("WS_BOT-%d", id)
}

func (wc *WSClient) ReadMess() (err error) {
	_, p, err := wc.conn.ReadMessage()
	if err != nil {
		return err
	}
	fmt.Printf("%s received from server: %s", wc.name, string(p))
	return nil
}

func (wc *WSClient) SendMess(json string) (err error) {
	sendJoin := []byte(json)

	fmt.Printf("%s send to server: %s \n", wc.name, json)
	err = wc.conn.WriteMessage(1, sendJoin)
	if err != nil {
		return err
	}

	return nil
}

func (wc *WSClient) CreateGame() (err error) {
	if err = wc.ReadMess(); err != nil {
		return err
	}
	fmt.Printf("%s try to create game \n", wc.name)
	if err = wc.SendMess(`{"action" : "CREATE_GAME"}`); err != nil {
		return err
	}
	if err = wc.ReadMess(); err != nil {
		return err
	}
	if err = wc.SendMess(fmt.Sprintf(`{"name" : "%s"}`, wc.name)); err != nil {
		return err
	}
	for {
		if err = wc.ReadMess(); err != nil {
			return err
		}
	}
}
