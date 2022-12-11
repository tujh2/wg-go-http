package api

import (
	"wg-go-http/model"
	"wg-go-http/wgutils"

	"gopkg.in/macaron.v1"
)

func addPeersHandler(ctx *macaron.Context, peersJson model.AddPeers) {
	err := wgutils.AddDevicePeers(peersJson.InterfaceName, peersJson.Peers)
	if err != nil {
		panic(err)
	}
}
