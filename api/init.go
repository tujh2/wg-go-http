package api

import (
	"wg-go-http/model"

	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
)

func InitApiRoutes(m *macaron.Macaron) {

	m.Post("/wgsetup", binding.Bind(model.JsonConfig{}), wgSetupHandler)
	m.Post("/addpeers", binding.Bind(model.AddPeers{}), addPeersHandler)
}
