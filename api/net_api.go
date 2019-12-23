package api

import (
	"github.com/idena-network/idena-go/ipfs"
	"github.com/idena-network/idena-go/protocol"
)

// NetApi offers helper utils
type NetApi struct {
	pm        *protocol.IdenaGossipHandler
	ipfsProxy ipfs.Proxy
}

// NewNetApi creates a new NetApi instance
func NewNetApi(pm *protocol.IdenaGossipHandler, ipfsProxy ipfs.Proxy) *NetApi {
	return &NetApi{pm, ipfsProxy}
}

func (api *NetApi) PeersCount() int {
	return api.pm.PeersCount()
}

type Peer struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RemoteAddr string `json:"addr"`
}

func (api *NetApi) Peers() []Peer {
	peers := make([]Peer, 0)
	return peers
}

func (api *NetApi) IpfsAddress() string {
	return api.pm.Endpoint()
}
