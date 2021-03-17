// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// Copyright 2020 ETH Zurich modifications to add support for SCION
package socket

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/lucas-clemente/quic-go"
	"github.com/netsec-ethz/scion-apps/pkg/appnet/appquic"
	"io"
	"net"
)

var _ net.Conn = &SingleStream{}

type SingleStream struct {
	quic.Stream
	session quic.Session
}

func (s SingleStream) LocalAddr() net.Addr {
	return s.session.LocalAddr()
}

func (s SingleStream) RemoteAddr() net.Addr {
	return s.session.RemoteAddr()
}

func DialAddr(remoteAddr string) (*SingleStream, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"scionftp"},
	}

	quicConfig := &quic.Config{
		KeepAlive: true,
	}

	session, err := appquic.Dial(remoteAddr, tlsConfig, quicConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to dial %s: %s", remoteAddr, err)
	}

	stream, err := session.OpenStream()
	if err != nil {
		return nil, err
	}
	err = sendHandshake(stream) // needed to unblock AcceptStream()
	if err != nil {
		return nil, err
	}

	return &SingleStream{stream, session}, nil
}

// Accept accepts a QUIC session with exactly one stream on listener.
// Note, that this function can not live at the listener due to hiding the session.
func Accept(listener quic.Listener) (*SingleStream, error) {
	session, err := listener.Accept(context.Background())
	if err != nil {
		return nil, err
	}

	stream, err := session.AcceptStream(context.Background())
	if err != nil {
		return nil, err
	}

	// AcceptStream() blocks until first data arrives, so we need to:
	err = consumeHandshake(stream)
	if err != nil {
		return nil, err
	}
	return &SingleStream{
		Stream:  stream,
		session: session,
	}, nil
}

func sendHandshake(rw io.ReadWriter) error {
	msg := []byte{200}
	_, err := rw.Write(msg)
	return err
}

func consumeHandshake(rw io.ReadWriter) error {
	msg := make([]byte, 1)
	n, err := rw.Read(msg)
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("invalid handshake received")
	}

	return nil
}
