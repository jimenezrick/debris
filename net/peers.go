package net

import "sync"

type void struct{}



type peers struct {
	mtx sync.RWMutex
	localAddr net.Addr
	conns map[string]peer
}




func Start() *peers {
	lis := startRpc()
	return &peers{
		make(map[string]bool),
		lis.Addr(),
	}
}






func (ps *peers) addPeer(peer string) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if _, ok := ps.connected[peer]; ok {
		return
	}
}

func (t *T) ConnectPeer(argType T1, replyType *T2) error




func (t *T) RemoteAddPeer(addr string, _ *void) error
