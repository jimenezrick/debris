package peer

import (
	"time"
	"net/rpc"
)

type void struct{}

type peer struct {
	addr string
	client rpc.Client
	beats chan bool
}

func newPeer(addr string) *peer {
	// TODO: hacer el Dial en sendHeartbeats()
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	peer := &peer{addr, client, make(chan bool, 1)}
	//go peer.sendHeartbeats()
	//go peer.monitorHeartbeats()
	return peer
}

func (p *peer) sendHeartbeats() {
	defer p.client.Close()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C:
		if err := p.client.Call("peers.Heartbeat", void{}, nil); err {
			panic(err)
		}
	}
}

func (p *peer) monitorHeartbeats() {
	for {
		<-p.c
		println("peer: heartbeat received")
	}
}

func (ps *peers) Heartbeat(addr string, _ *void) error {
	ps.mtx.RLock()
	defer ps.mtx.RUnlock()
	peer, ok := ps.conns[addr]
	if !ok {
		panic("Heartbeat: unknown peer")
	}

	peer.beats <- true
	return nil
}

// TODO: removePeer()
