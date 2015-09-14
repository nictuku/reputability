// Definitions:
// - node: routing element. Does not have intrinsic reliability or reputation
// - peer: service element (e.g: storage server). Has reputation.

// Procedures:
// - Find most reliable peer with [search criteria]
// - Find a node's total reliability score, according to direct nodes.
//
// Upkeep:
// - Recalculate neighbor's reliability
//
// Basic operation:
// - initial node layout A - B - C - E
// - E needs storage. Finds A
// - A stores file for E
// - E now trusts A.
// - E reliability score for A = bytes * seconds
// - A reliability score for E = 0
//  - C
// Example procedure use cases:
// - Find my own reputation and display it in the UI.
//
// Basic principles: - node A keeps track of other nodes, with two attributes: -
// node should not get penalized from opening new connections (should have an
// incentive, maybe use sum of reputation)

// Package reputation provides a software simulation of the reputability system
// of distributed reliability information.
package reputation

import (
	"net"
	"time"
)

type serviceClient interface {
	isReliable(*peer) bool
}

// peer is an entity that agreed to provide a service.
type peer struct {
	addr net.Addr
	time time.Time
}

func (p *peer) renewReputation(service serviceClient) {
	if !service.isReliable(p) {
		p.time = time.Time{}
		return
	}
	if p.time.IsZero() {
		p.time = time.Now()
	}
}

type node struct {
	addr net.Addr
}

type agent struct {
	service serviceClient
	nodes   []*node
	peers   []*peer
}

func (n *agent) AddNode(addr net.Addr) {
	n.nodes = append(n.nodes, &node{addr})
}

func (n *agent) Reputation(addr net.Addr) int64 {
	// TODO: Use a logistic function to grow the score until it basically stops
	// growing after 15 days or so.
	for _ = range n.nodes {

	}
	return 0
}

func (n *agent) Search(query string) ([]*node, error) {
	// TODO
	ret := []*node{}
	for _, nd := range n.nodes {
		if rpc.hasSpace(nd) {
			ret = append(ret, nd)
		}
	}
	return ret, nil
}

func (n *agent) upkeep() {
	for _, p := range n.peers {
		p.renewReputation(n.service)
	}
}
