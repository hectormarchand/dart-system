// Package sse implements a minimal Server-Sent Events broker on top of
// net/http, used to push live game and dart-hit updates to connected clients.
package sse

import (
	"fmt"
	"net/http"
	"sync"
)

// Broker fans out published messages to every currently subscribed client.
type Broker struct {
	mu      sync.Mutex
	clients map[chan string]struct{}
}

func NewBroker() *Broker {
	return &Broker{clients: make(map[chan string]struct{})}
}

func (b *Broker) subscribe() chan string {
	ch := make(chan string, 8)
	b.mu.Lock()
	b.clients[ch] = struct{}{}
	b.mu.Unlock()
	return ch
}

func (b *Broker) unsubscribe(ch chan string) {
	b.mu.Lock()
	delete(b.clients, ch)
	b.mu.Unlock()
	close(ch)
}

// Publish sends data to every subscribed client. A slow/stalled client is
// dropped rather than blocking the publisher.
func (b *Broker) Publish(data string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for ch := range b.clients {
		select {
		case ch <- data:
		default:
		}
	}
}

// Serve upgrades the response to an SSE stream, sends `initial` as the first
// event, then relays every subsequent Broker.Publish call until the client
// disconnects.
func Serve(w http.ResponseWriter, r *http.Request, broker *Broker, initial string) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	ch := broker.subscribe()
	defer broker.unsubscribe(ch)

	fmt.Fprintf(w, "data: %s\n\n", initial)
	flusher.Flush()

	for {
		select {
		case data, ok := <-ch:
			if !ok {
				return
			}
			fmt.Fprintf(w, "data: %s\n\n", data)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
