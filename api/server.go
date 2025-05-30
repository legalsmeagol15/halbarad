package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // For demo purposes; restrict in production!
	},
}

// Client represents a connected websocket client
type Client struct {
	conn  *websocket.Conn
	send  chan []byte
	token token
}

type token string

// Hub maintains active clients and broadcasts messages to them
type Hub struct {
	clients    map[*Client]token
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]token),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = ""
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.mu.Unlock()
		case message := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					delete(h.clients, client)
					close(client.send)
				}
			}
			h.mu.Unlock()
		}
	}
}

func serveSession(hub *Hub, w http.ResponseWriter, r *http.Request) {

	// Authenticate
	entry := registerEntry{}
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		// Handle decoding error
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	// Upgrade to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	// Put the client on the active hub
	client := &Client{conn: ws, send: make(chan []byte, 256)}
	hub.register <- client

	// Read and write buffers
	go client.writePump()
	client.readPump(hub)
}

func (c *Client) readPump(hub *Hub) {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		// Bounce messages to all clients (or handle individually)
		hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
	c.conn.Close()
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	creds := credentials{}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		// Handle decoding error
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad account credentials"))
	} else if creds.Password, err = hashPswd(creds.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error while creating account"))
	} else if _, err = getClient(creds); err == nil {
		// Would indicate a duplicate
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad account credentials"))
	} else if err = createAccount(creds); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error while creating account"))
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func handleDemo(w http.ResponseWriter, r *http.Request) {
	// session, _ := store.Get(r, "session-name")
	// session.Values["authenticated"] = true // or set a demo user id
	// session.Save(r, w)

	// Respond with an appropriate message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged in as demo user"))
}

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {
		serveSession(hub, w, r)
	})
	http.HandleFunc("/create_account", handleCreateAccount)
	http.HandleFunc("/logged_in", handleDemo)

	log.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
