package reputation

import "testing"

func TestReputation(t *testing.T) {

	storage := newRemoteStorage()
	testClient := &agent{
		service: storage,
	}

	fake := newFakeNet()
	data := []byte("FOOOO")

	// Create simulated nodes and connect them to the test client.
	nodes := []*fakeNode{
		&fakeNode{"A", false},
		&fakeNode{"B", true}, // has space
		&fakeNode{"C", false},
	}
	for _, n := range nodes {
		fake.CreateNode(n, n.hasStorageSpace)
		testClient.AddNode(n)
	}

	// Node B has available space.

	// Find peer.
	candidatePeers, _ := testClient.Search("query")

	// TODO: negotiate.
	preferred := &peer{addr: candidatePeers[0].addr}

	// Use peer service.
	storage.storeAt(data, preferred)

	// reputation.
	preferred.renewReputation(storage)
	t.Logf("peer %q looking good since: %v", preferred.addr.String(), preferred.time.String())

	testClient.upkeep()

	// Verify that when you ask neighbor nodes, they think our peer is awesome.
}
