package model

type JsonConfig struct {
	InterfaceName string     `json:"interfaceName"`
	PrivateKey    string     `json:"privateKey"`
	BaseAddr      string     `json:"baseAddr"`
	ListenPort    int        `json:"listenPort"`
	Peers         []JsonPeer `json:"peers"`
}

type AddPeers struct {
	InterfaceName string     `json:"interfaceName"`
	Peers         []JsonPeer `json:"peers"`
}

type JsonPeer struct {
	PeerPublicKey string `json:"peerPublicKey"`
	PeerIP        string `json:"peerIp"`
}
