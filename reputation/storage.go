package reputation

import "fmt"

func newRemoteStorage() *remoteStorage {
	return &remoteStorage{
		objects: make(map[string][]byte),
	}
}

type remoteStorage struct {
	objects map[string][]byte
}

func (r *remoteStorage) storeAt(object []byte, p *peer) error {
	r.objects[p.addr.String()] = object // keep local copy. XXX: store chunk hashes.
	return rpc.write(object, p)
}

func (r *remoteStorage) probePeer(object []byte, p *peer) error {
	if rpc.hasStored(object, p) {
		return nil
	}
	return fmt.Errorf("object not found in storage")
}

func (r *remoteStorage) isReliable(p *peer) bool {
	// did I ever ask p to store objects for me?
	obj, ok := r.objects[p.addr.String()]
	if !ok {
		return false
	}
	// does it still have it?
	err := r.probePeer(obj, p)
	return err == nil
}
