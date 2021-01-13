// Copyright 2020 ETH Zurich
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package socket

import (
	"encoding/binary"
	"fmt"

	"github.com/netsec-ethz/scion-apps/internal/ftp/striping"
)

// A ReadWorker should be dispatched and runs until it
// receives the closing connection
// Does not need to be closed since it's closed
// automatically
type ReadWorker struct {
	socket DataSocket
	// ctx    context.Context // Currently unused
}

func NewReadWorker(socket DataSocket) *ReadWorker {
	return &ReadWorker{socket: socket}
}

// Keeps running until it receives an EOD flag
func (s *ReadWorker) Run(push chan<- *striping.Segment) {
	for {
		seg, err := receiveNextSegment(s.socket)
		if err != nil {
			fmt.Printf("Failed to receive segment: %s\n", err)
		}

		push <- seg

		if seg.ContainsFlag(striping.BlockFlagEndOfData) {
			return
		}

	}
}

func receiveNextSegment(socket DataSocket) (*striping.Segment, error) {
	header := &striping.Header{}
	err := binary.Read(socket, binary.BigEndian, header)
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %s", err)
	}

	data := make([]byte, header.ByteCount)
	cur := 0

	// Read all bytes
	for {
		n, err := socket.Read(data[cur:header.ByteCount])
		if err != nil {
			return nil, fmt.Errorf("failed to read payload: %s", err)
		}

		cur += n
		if cur == int(header.ByteCount) {
			return striping.NewSegmentWithHeader(header, data), nil
		}
	}
}
