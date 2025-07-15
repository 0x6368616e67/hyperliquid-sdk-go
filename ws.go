package hyperliquid

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type wsConn struct {
	url           string
	conn          *websocket.Conn
	wsLock        sync.Mutex
	topics        []any
	closeChan     chan struct{}
	safecloseChan chan struct{}
	pingTicker    *time.Ticker

	MessageChan chan []byte
}

func newWSConn(endpoint string) *wsConn {
	conn := &wsConn{
		closeChan:     make(chan struct{}),
		safecloseChan: make(chan struct{}),
	}
	wsConn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		return nil
	}
	conn.conn = wsConn

	conn.conn.SetReadLimit(WSMaxMessageSize)
	conn.conn.SetReadDeadline(time.Now().Add(WSPongWait))
	conn.conn.SetPongHandler(func(string) error {
		conn.conn.SetReadDeadline(time.Now().Add(WSPongWait))
		return nil
	})
	return conn
}

func (c *wsConn) writeJSON(message any) (err error) {
	c.wsLock.Lock()
	defer c.wsLock.Unlock()
	return c.conn.WriteJSON(message)
}

func (c *wsConn) Connect() (err error) {
	c.ping()
	go c.loop()
	return
}

func (c *wsConn) Subscribe(message any) (err error) {
	for _, t := range c.topics {
		if t == message {
			return nil
		}
	}
	msg := struct {
		Method       string `json:"method"`
		Subscription any    `json:"subscription"`
	}{
		Method:       "subscribe",
		Subscription: message,
	}
	c.topics = append(c.topics, message)
	return c.writeJSON(msg)
}

func (c *wsConn) Close() {
	if c.pingTicker != nil {
		c.pingTicker.Stop()
	}

	c.wsLock.Lock()
	defer c.wsLock.Unlock()
	<-c.safecloseChan
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
}

func (c *wsConn) loop() {
	c.pingTicker = time.NewTicker(WSHeartbeatInterval)

	for {
		select {
		case <-c.pingTicker.C:
			if err := c.ping(); err != nil {
				// retry 3 timeo
				log.Fatalf("ping error:%v", err)
				return
			}
		case <-c.closeChan:
			c.pingTicker.Stop()
			c.safecloseChan <- struct{}{}
			log.Panicln("close")
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					// normal close
				} else if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure) {
					// closed with erro
				} else {
					// read error
				}
				break
			}
			fmt.Printf("got message:%+v", string(message))
			//c.MessageChan <- message
		}
	}
}

func (c *wsConn) ping() (err error) {
	c.wsLock.Lock()
	err = c.conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(3*time.Second))
	c.wsLock.Unlock()

	if err != nil {
		return
	}
	return
}
