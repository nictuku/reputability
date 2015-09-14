package reputation

import (
	"crypto/md5"
	"fmt"
	"net"
)

type fakeNode struct {
	addr            string
	hasStorageSpace bool
}

func (f *fakeNode) String() string {
	return f.addr
}

func (f *fakeNode) Network() string {
	return "fake"
}

func newFakeNet() *fakeNet {
	return &fakeNet{
		nodes: make(map[net.Addr]bool),
	}
}

type fakeNet struct {
	nodes map[net.Addr]bool
}

func (fk *fakeNet) CreateNode(n net.Addr, hasStorageSpace bool) {
	fk.nodes[n] = true
	rpc.nodeHasSpace[n] = hasStorageSpace
}

type fakeRPC struct {
	m            map[string]bool
	nodeHasSpace map[net.Addr]bool
}

func (r *fakeRPC) write(object []byte, p *peer) error {

	h := fmt.Sprintf("%x", md5.Sum(object))
	r.m[h+"@"+p.addr.String()] = true
	return nil
}

func (r *fakeRPC) hasStored(object []byte, p *peer) bool {
	h := fmt.Sprintf("%x", md5.Sum(object))
	_, ok := r.m[h+"@"+p.addr.String()]
	return ok
}

func (r *fakeRPC) hasSpace(n *node) bool {
	return r.nodeHasSpace[n.addr]
}

var rpc = fakeRPC{
	m:            make(map[string]bool),
	nodeHasSpace: make(map[net.Addr]bool),
}
