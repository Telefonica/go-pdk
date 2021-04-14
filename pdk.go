/*
Package Telefonica/go-pdk implements Kong's Plugin Development Kit for Go.

It directly parallels the existing kong PDK for Lua plugins.

Kong plugins written in Go implement event handlers as methods on the Plugin's
structure, with the given signature:

	func (conf *MyConfig) Access (kong *pdk.PDK) {
		...
	}

The `kong` argument of type `*pdk.PDK` is the entrypoint for all PDK functions.
For example, to get the client's IP address, you'd use `kong.Client.GetIp()`.
*/
package pdk

import (
	"github.com/Telefonica/go-pdk/client"
	"github.com/Telefonica/go-pdk/ctx"
	"github.com/Telefonica/go-pdk/ip"
	"github.com/Telefonica/go-pdk/log"
	"github.com/Telefonica/go-pdk/nginx"
	"github.com/Telefonica/go-pdk/node"
	"github.com/Telefonica/go-pdk/request"
	"github.com/Telefonica/go-pdk/response"
	"github.com/Telefonica/go-pdk/router"
	"github.com/Telefonica/go-pdk/service"
	service_request "github.com/Telefonica/go-pdk/service/request"
	service_response "github.com/Telefonica/go-pdk/service/response"
)

// PDK go pdk module
type PDK struct {
	Client          client.Client
	Ctx             ctx.Ctx
	Log             log.Log
	Nginx           nginx.Nginx
	Request         request.Request
	Response        response.Response
	Router          router.Router
	IP              ip.Ip
	Node            node.Node
	Service         service.Service
	ServiceRequest  service_request.Request
	ServiceResponse service_response.Response
}

// Init initialize go pdk.  Called by the pluginserver at initialization.
func Init(ch chan interface{}) *PDK {
	return &PDK{
		Client:          client.New(ch),
		Ctx:             ctx.New(ch),
		Log:             log.New(ch),
		Nginx:           nginx.New(ch),
		Request:         request.New(ch),
		Response:        response.New(ch),
		Router:          router.New(ch),
		IP:              ip.New(ch),
		Node:            node.New(ch),
		Service:         service.New(ch),
		ServiceRequest:  service_request.New(ch),
		ServiceResponse: service_response.New(ch),
	}
}
