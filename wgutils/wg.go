package wgutils

import (
	"net"
	"wg-go-http/model"

	"github.com/vishvananda/netlink"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

const WireguardInterfaceType = "wireguard"

type WireguardLink struct {
	LinkAttrs netlink.LinkAttrs
}

func (device *WireguardLink) Attrs() *netlink.LinkAttrs {
	return &device.LinkAttrs
}

func (device *WireguardLink) Type() string {
	return WireguardInterfaceType
}

func GenerateWgConfig(jsonCofig model.JsonConfig) (wgtypes.Config, error) {
	privkey, err := wgtypes.ParseKey(jsonCofig.PrivateKey)
	if err != nil {
		return wgtypes.Config{}, err
	}

	config := wgtypes.Config{
		PrivateKey:   &privkey,
		ListenPort:   &jsonCofig.ListenPort,
		FirewallMark: nil,
		ReplacePeers: true,
		Peers:        []wgtypes.PeerConfig{},
	}

	for _, peer := range jsonCofig.Peers {

		pubkey, err := wgtypes.ParseKey(peer.PeerPublicKey)
		if err != nil {
			return config, err
		}

		allowedIP := net.ParseIP(peer.PeerIP)

		config.Peers = append(config.Peers, wgtypes.PeerConfig{
			PublicKey:         pubkey,
			Remove:            false,
			UpdateOnly:        false,
			ReplaceAllowedIPs: true,
			AllowedIPs: append([]net.IPNet{}, net.IPNet{
				IP:   allowedIP,
				Mask: net.IPv4Mask(255, 255, 255, 255),
			}),
		})
	}

	return config, nil
}

func AddDevicePeers(name string, peers []model.JsonPeer) error {
	client, err := wgctrl.New()
	if err != nil {
		return err
	}
	defer client.Close()
	dev, err := client.Device(name)
	if err != nil {
		return err
	}

	config := wgtypes.Config{
		PrivateKey:   &dev.PrivateKey,
		ListenPort:   &dev.ListenPort,
		FirewallMark: nil,
		ReplacePeers: false,
		Peers:        []wgtypes.PeerConfig{},
	}

	for _, peer := range peers {

		pubkey, err := wgtypes.ParseKey(peer.PeerPublicKey)
		if err != nil {
			return err
		}

		allowedIP := net.ParseIP(peer.PeerIP)

		config.Peers = append(config.Peers, wgtypes.PeerConfig{
			PublicKey:         pubkey,
			Remove:            false,
			UpdateOnly:        false,
			ReplaceAllowedIPs: true,
			AllowedIPs: append([]net.IPNet{}, net.IPNet{
				IP:   allowedIP,
				Mask: net.IPv4Mask(255, 255, 255, 255),
			}),
		})
	}

	client.ConfigureDevice(name, config)

	if err != nil {
		return err
	}
	return nil
}

func ApplyWgConfig(name string, config wgtypes.Config) error {
	client, err := wgctrl.New()
	if err != nil {
		return err
	}
	defer client.Close()
	err = client.ConfigureDevice(name, config)
	if err != nil {
		return err
	}
	return nil
}
