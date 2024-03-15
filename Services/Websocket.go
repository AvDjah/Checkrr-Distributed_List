package Services

type ClientHub struct {
	clients   map[*Client]bool
	broadcast chan []byte
}

type Client struct {
}
