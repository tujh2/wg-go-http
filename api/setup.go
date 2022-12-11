package api

import (
	"fmt"
	"os/exec"
	"reflect"
	"wg-go-http/model"
	"wg-go-http/wgutils"

	"github.com/vishvananda/netlink"
	"gopkg.in/macaron.v1"
)

const postUpScript = "postUp.sh"

func wgSetupHandler(ctx *macaron.Context, jsonConfig model.JsonConfig) {

	link, err := netlink.LinkByName(jsonConfig.InterfaceName)

	recreate := true
	if err != nil {
		if reflect.TypeOf(err) == reflect.TypeOf(netlink.LinkNotFoundError{}) {
			attrs := netlink.NewLinkAttrs()
			attrs.Name = jsonConfig.InterfaceName
			adderr := netlink.LinkAdd(&wgutils.WireguardLink{LinkAttrs: attrs})

			if adderr != nil {
				panic(adderr)
			}

			link, err = netlink.LinkByName(jsonConfig.InterfaceName)
			if err != nil {
				panic(err)
			}
			recreate = false
		} else {
			panic(err)
		}
	}

	if recreate {
		if link.Type() != wgutils.WireguardInterfaceType {
			panic(fmt.Errorf("INTERFACE %s EXIST AND IS NOT WIREGUARD TYPE", jsonConfig.InterfaceName))
		}
		err := netlink.LinkDel(link)
		if err != nil {
			panic(err)
		}
		attrs := netlink.NewLinkAttrs()
		attrs.Name = jsonConfig.InterfaceName
		adderr := netlink.LinkAdd(&wgutils.WireguardLink{LinkAttrs: attrs})

		if adderr != nil {
			panic(adderr)
		}

		link, err = netlink.LinkByName(jsonConfig.InterfaceName)
		recreate = false
		if err != nil {
			panic(err)
		}
	}

	err = netlink.LinkSetUp(link)

	if err != nil {
		panic(err)
	}

	baseAddr, err := netlink.ParseAddr(jsonConfig.BaseAddr)

	if err != nil {
		panic(err)
	}

	err = netlink.AddrAdd(link, baseAddr)
	if err != nil {
		panic(err)
	}

	wgconfig, err := wgutils.GenerateWgConfig(jsonConfig)
	if err != nil {
		panic(err)
	}

	err = wgutils.ApplyWgConfig(jsonConfig.InterfaceName, wgconfig)
	if err != nil {
		panic(err)
	}

	exec.Command("/bin/sh", postUpScript)

}
