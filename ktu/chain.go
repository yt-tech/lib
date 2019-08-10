// Copyright 2017 Tim Oster. All rights reserved.
// Use of this source code is governed by the MIT license.
// More information can be found in the LICENSE file.

package ktu

import "sync"

type chain struct {
	next      orderNumber
	start     *chainLink
	length    byte
	maxLength byte
	mutex     sync.Mutex
}

type chainLink struct {
	next   *chainLink
	packet *packet
}

func newChain(maxLength byte) *chain {
	return &chain{
		maxLength: maxLength,
	}
}

func (chain *chain) reset() {
	chain.mutex.Lock()
	chain.next = 0
	chain.start = nil
	chain.length = 0
	chain.mutex.Unlock()
}

func (chain *chain) chain(packet *packet) {
	chain.mutex.Lock()
	if chain.start == nil {
		chain.start = &chainLink{next: nil, packet: packet}
	} else {
		var link *chainLink

		for l := chain.start; l != nil; l = l.next {
			if greaterThanOrder(packet.order, l.packet.order) {
				link = l
			} else {
				break
			}
		}

		if link == nil {
			chain.start = &chainLink{next: chain.start, packet: packet}
		} else {
			link.next = &chainLink{next: link.next, packet: packet}
		}
	}

	if chain.length >= chain.maxLength {
		chain.start = chain.start.next
		chain.length--
	}

	chain.length++
	chain.mutex.Unlock()
}

func (chain *chain) popConsecutive() *chainLink {
	chain.mutex.Lock()
	defer chain.mutex.Unlock()

	var last *chainLink

	for l := chain.start; l != nil; l = l.next {
		if l.packet.order == chain.next {
			chain.length--
			chain.next++
			last = l
		} else {
			break
		}
	}

	if last != nil {
		start := chain.start
		chain.start = last.next
		last.next = nil
		return start
	}

	return nil
}

func (chain *chain) skip() {
	chain.mutex.Lock()
	if chain.start != nil {
		chain.next = chain.start.packet.order
	}
	chain.mutex.Unlock()
}
