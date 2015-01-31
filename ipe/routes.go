// Copyright 2014 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ipe

import (
	"net/http"
)

// A route
type Route struct {
	Name             string
	Method           string
	Pattern          string
	HandlerFunc      http.HandlerFunc
	RequiresRestAuth bool
}

type Routes []Route

var routes = Routes{
	Route{
		"PostEvents",
		"POST",
		"/apps/{app_id}/events",
		postEvents,
		true,
	},
	Route{
		"GetChannels",
		"GET",
		"/apps/{app_id}/channels",
		getChannels,
		true,
	},
	Route{
		"GetChannel",
		"GET",
		"/apps/{app_id}/channels/{channel_name}",
		getChannel,
		true,
	},
	Route{
		"GetChannelUsers",
		"GET",
		"/apps/{app_id}/channels/{channel_name}/users",
		getChannelUsers,
		true,
	},
	Route{
		"Websocket",
		"GET",
		"/app/{key}",
		wsHandler,
		false,
	},
}
