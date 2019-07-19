package socket

import (
	"io"
	"time"
)

type Socket interface {
	io.Reader
	io.Writer
	io.Closer
}

type MultiSocket struct {
	*ReaderSocket
	*WriterSocket
	closed bool
}

func (m *MultiSocket) SetDeadline(t time.Time) error {
	panic("implement me")
}

var _ DataSocket = &MultiSocket{}

func (MultiSocket) Host() string {
	return "hostaddress"
}

func (MultiSocket) Port() int {
	return 9999
}

// Only the client should close the socket
// Sends the closing message
func (m *MultiSocket) Close() error {
	return m.WriterSocket.Close()
}

var _ DataSocket = &MultiSocket{}

func NewMultiSocket(sockets []DataSocket, maxLength int) *MultiSocket {
	return &MultiSocket{
		NewReadsocket(sockets),
		NewWriterSocket(sockets, maxLength),
		false,
	}
}
