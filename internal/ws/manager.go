package ws

import "sync"

type Manager struct {
	mu      sync.RWMutex
	clients map[uint64]*Client
}

func NewManager() *Manager {
	return &Manager{
		clients: make(map[uint64]*Client),
	}
}

func (m *Manager) Register(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[client.UserID] = client
}

func (m *Manager) Unregister(userID uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, userID)
}
