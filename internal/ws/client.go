package ws

type Client struct {
	UserID uint64
	Send   chan []byte
}

func NewClient(userID uint64) *Client {
	return &Client{
		UserID: userID,
		Send:   make(chan []byte, 256),
	}
}
