// Copyright 2017 Tim Oster. All rights reserved.
// Use of this source code is governed by the MIT license.
// More information can be found in the LICENSE file.

package ktu

import "sync"

type sequenceBuffer struct {
	size      sequenceNumber
	sequences []sequenceNumber
	states    []bool
	mutex     sync.Mutex
}

func newSequenceBuffer(size sequenceNumber) *sequenceBuffer {
	buffer := new(sequenceBuffer)
	buffer.size = size
	buffer.sequences = make([]sequenceNumber, size)
	buffer.states = make([]bool, size)
	return buffer
}

func (buffer *sequenceBuffer) reset() {
	buffer.mutex.Lock()

	for i := sequenceNumber(0); i < buffer.size; i++ {
		buffer.sequences[i] = 0
		buffer.states[i] = false
	}
	buffer.mutex.Unlock()
}

func (buffer *sequenceBuffer) get(sequence sequenceNumber) bool {
	buffer.mutex.Lock()
	defer buffer.mutex.Unlock()

	if buffer.sequences[sequence%buffer.size] != sequence {
		return false
	}

	return buffer.states[sequence%buffer.size]
}

func (buffer *sequenceBuffer) set(sequence sequenceNumber, value bool) {
	buffer.mutex.Lock()
	buffer.sequences[sequence%buffer.size] = sequence
	buffer.states[sequence%buffer.size] = value
	buffer.mutex.Unlock()
}
